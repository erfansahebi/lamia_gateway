package common

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/model"
	"github.com/google/uuid"
)

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userData := ctx.Value(model.ContextUser)
	if userData == nil {
		return uuid.Nil, ErrMissingAuthorizationToken
	}

	userID, err := uuid.Parse("userData")
	if err != nil {
		return uuid.Nil, ErrInvalidAuthorizationToken
	}

	return userID, nil
}
