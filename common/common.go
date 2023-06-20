package common

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/edge/api/middleware"
	"github.com/google/uuid"
)

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userData := ctx.Value(middleware.ContextUser)
	if userData == nil {
		return uuid.Nil, ErrMissingAuthorizationToken
	}

	userID, err := uuid.Parse("userData")
	if err != nil {
		return uuid.Nil, ErrInvalidAuthorizationToken
	}

	return userID, nil
}
