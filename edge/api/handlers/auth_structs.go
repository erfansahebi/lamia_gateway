package handlers

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/di"
)

type AuthorizationUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// Register

type RegisterRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (rr *RegisterRequest) Validate(ctx context.Context, container di.DIContainerInterface) error {
	switch {
	case rr.FirstName == "":
		return ErrEmptyFields
	case rr.LastName == "":
		return ErrEmptyFields
	case rr.Email == "":
		return ErrEmptyFields
	case rr.Password == "":
		return ErrEmptyFields
	case rr.PasswordConfirm == "":
		return ErrEmptyFields
	case rr.Password != rr.PasswordConfirm:
		return ErrPasswordMatch
	}

	return nil
}

type RegisterResponse struct {
	User               AuthorizationUserResponse `json:"user"`
	AuthorizationToken string                    `json:"authorization_token"`
}

// Login

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lr *LoginRequest) Validate(ctx context.Context, container di.DIContainerInterface) error {
	switch {
	case lr.Email == "":
		return ErrEmptyFields
	case lr.Password == "":
		return ErrEmptyFields
	}

	return nil
}

type LoginResponse struct {
	User               AuthorizationUserResponse `json:"user"`
	AuthorizationToken string                    `json:"authorization_token"`
}
