package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/erfansahebi/lamia_gateway/di"
	"github.com/erfansahebi/lamia_shared/go/log"
	"net/http"
	"strings"
)

type Handler struct {
	AppCtx context.Context
	Di     di.DIContainerInterface
}

type RequestHandler func(r *http.Request) (interface{}, int, error)

type ResponseStruct struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

func Response(writer http.ResponseWriter, detail ResponseStruct, code int) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(code)

	json.NewEncoder(writer).Encode(detail)
}

func HTTPErr(w http.ResponseWriter, statusCode int, errs ...error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	detail := map[string]interface{}{
		"error": errs[0].Error(),
	}

	if len(errs) > 1 {
		detail["errors"] = errs[1:]
	}

	json.NewEncoder(w).Encode(detail)
}

type APIResource interface {
	Expose() interface{}
}

type APIErr interface {
	error
	Status() int
}

func Wrap(rq RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, status, err := rq(r)

		if err != nil {
			log.WithError(err).Errorf(r.Context(), "http request error")
			HTTPErr(w, status, err)
			return
		}

		if response, ok := response.(APIResource); ok {
			Response(w, ResponseStruct{Data: response.Expose()}, status)
			return
		}

		Response(w, ResponseStruct{Data: response}, status)

	}
}

func HandleErrorFromGrpc(grpcError error) error {
	splitError := strings.Split(grpcError.Error(), "= ")
	return errors.New(splitError[len(splitError)-1])
}
