package common

import "errors"

var (
	ErrMissingAuthorizationToken    = errors.New("missing authorization token")
	ErrInvalidAuthorizationToken    = errors.New("invalid authorization token")
	ErrUnAuthorized                 = errors.New("unAuthorized")
	GenericHTTPServerErr            = errors.New("the server has failed to process the request, please do not retry")
	ErrAccessDenied                 = errors.New("access denied")
	ErrEmptyFields                  = errors.New("please fill the fields")
	ErrPasswordMatch                = errors.New("passwords did not match")
	ErrIncorrectCountryCode         = errors.New("incorrect country code")
	ErrFailedToReadRequestBody      = errors.New("failed to read request body")
	ErrFailedToUnmarshalRequestBody = errors.New("failed to unmarshal request body")
	ErrInternalError                = errors.New("internal error")
	ErrInvalidURL                   = errors.New("GetHttpUrl is formatted improperly")
	ErrPermissionDenied             = errors.New("permission denied")
	ErrInvalidDate                  = errors.New("invalid date format")
)
