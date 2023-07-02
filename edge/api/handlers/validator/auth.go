package validator

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/common"
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
		return common.ErrEmptyFields
	case rr.LastName == "":
		return common.ErrEmptyFields
	case rr.Email == "":
		return common.ErrEmptyFields
	case rr.Password == "":
		return common.ErrEmptyFields
	case rr.PasswordConfirm == "":
		return common.ErrEmptyFields
	case rr.Password != rr.PasswordConfirm:
		return common.ErrPasswordMatch
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
		return common.ErrEmptyFields
	case lr.Password == "":
		return common.ErrEmptyFields
	}

	return nil
}

type LoginResponse struct {
	User               AuthorizationUserResponse `json:"user"`
	AuthorizationToken string                    `json:"authorization_token"`
}
