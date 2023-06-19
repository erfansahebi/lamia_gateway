package handlers

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/di"
)

type RegisterStruct struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (rs *RegisterStruct) Validate(ctx context.Context, container di.DIContainerInterface) error {
	switch {
	case rs.FirstName == "":
		return ErrEmptyFields
	case rs.LastName == "":
		return ErrEmptyFields
	case rs.Email == "":
		return ErrEmptyFields
	case rs.Password == "":
		return ErrEmptyFields
	}

	return nil
}

type LoginStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ls *LoginStruct) Validate(ctx context.Context, container di.DIContainerInterface) error {
	switch {
	case ls.Email == "":
		return ErrEmptyFields
	case ls.Password == "":
		return ErrEmptyFields
	}

	return nil
}
