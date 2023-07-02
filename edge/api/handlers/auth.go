package handlers

import (
	"encoding/json"
	"github.com/erfansahebi/lamia_gateway/common"
	"github.com/erfansahebi/lamia_gateway/edge/api/handlers/validator"
	authProto "github.com/erfansahebi/lamia_shared/go/proto/auth"
	"io"
	"net/http"
)

func (h *Handler) Register(r *http.Request) (interface{}, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrFailedToReadRequestBody
	}

	var pendData validator.RegisterRequest
	if err = json.Unmarshal(body, &pendData); err != nil {
		return nil, http.StatusBadRequest, common.ErrFailedToUnmarshalRequestBody
	}

	if err = pendData.Validate(r.Context(), h.Di); err != nil {
		return nil, http.StatusBadRequest, err
	}

	authenticateData, err := h.Di.Services().Auth().Client().Register(r.Context(), &authProto.RegisterRequest{
		User: &authProto.UserStruct{
			Id:        "",
			FirstName: pendData.FirstName,
			LastName:  pendData.LastName,
			Email:     pendData.Email,
			Password:  pendData.Password,
		},
	})
	if err != nil {
		return nil, http.StatusBadRequest, HandleErrorFromGrpc(err)
	}

	return validator.RegisterResponse{
		User: validator.AuthorizationUserResponse{
			ID:        authenticateData.User.Id,
			FirstName: authenticateData.User.FirstName,
			LastName:  authenticateData.User.LastName,
			Email:     authenticateData.User.Email,
		},
		AuthorizationToken: authenticateData.AuthorizationToken,
	}, http.StatusOK, nil
}

func (h *Handler) Login(r *http.Request) (interface{}, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, common.ErrFailedToReadRequestBody
	}

	var pendData validator.LoginRequest
	if err = json.Unmarshal(body, &pendData); err != nil {
		return nil, http.StatusBadRequest, common.ErrFailedToUnmarshalRequestBody
	}

	if err = pendData.Validate(r.Context(), h.Di); err != nil {
		return nil, http.StatusBadRequest, err
	}

	authenticateData, err := h.Di.Services().Auth().Client().Login(r.Context(), &authProto.LoginRequest{
		Email:    pendData.Email,
		Password: pendData.Password,
	})
	if err != nil {
		return nil, http.StatusBadRequest, HandleErrorFromGrpc(err)
	}

	return validator.LoginResponse{
		User: validator.AuthorizationUserResponse{
			ID:        authenticateData.User.Id,
			FirstName: authenticateData.User.FirstName,
			LastName:  authenticateData.User.LastName,
			Email:     authenticateData.User.Email,
		},
		AuthorizationToken: authenticateData.AuthorizationToken,
	}, http.StatusOK, nil
}

func (h *Handler) Logout(r *http.Request) (interface{}, int, error) {
	authorizationToken := r.Header.Get("Authorization")

	if _, err := h.Di.Services().Auth().Client().Logout(r.Context(), &authProto.LogoutRequest{
		AuthorizationToken: authorizationToken,
	}); err != nil {
		return nil, http.StatusBadRequest, HandleErrorFromGrpc(err)
	}

	return nil, http.StatusOK, nil
}
