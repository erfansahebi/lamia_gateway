package common

import "errors"

var (
	ErrMissingAuthorizationToken = errors.New("missing authorization token")
	ErrInvalidAuthorizationToken = errors.New("invalid authorization token")
	ErrUnAuthorized              = errors.New("unAuthorized")
)
