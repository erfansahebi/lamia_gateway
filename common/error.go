package common

import "errors"

var (
	ErrMissingAuthorizationToken = errors.New("missing authorization token")
	ErrUnAuthorized              = errors.New("unAuthorized")
)
