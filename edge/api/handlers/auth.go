package handlers

import (
	"encoding/json"
	authProto "github.com/erfansahebi/lamia_shared/services/auth"
	"io"
	"net/http"
)

func (h *Handler) Register(r *http.Request) (interface{}, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, ErrFailedToReadRequestBody
	}

	var pendData RegisterRequest
	if err = json.Unmarshal(body, &pendData); err != nil {
		return nil, http.StatusBadRequest, ErrFailedToUnmarshalRequestBody
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

	return RegisterResponse{
		User: AuthorizationUserResponse{
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
		return nil, http.StatusBadRequest, ErrFailedToReadRequestBody
	}

	var pendData LoginRequest
	if err = json.Unmarshal(body, &pendData); err != nil {
		return nil, http.StatusBadRequest, ErrFailedToUnmarshalRequestBody
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

	return LoginResponse{
		User: AuthorizationUserResponse{
			ID:        authenticateData.User.Id,
			FirstName: authenticateData.User.FirstName,
			LastName:  authenticateData.User.LastName,
			Email:     authenticateData.User.Email,
		},
		AuthorizationToken: authenticateData.AuthorizationToken,
	}, http.StatusOK, nil
}
