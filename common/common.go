package common

import (
	"context"
	"fmt"
	"github.com/erfansahebi/lamia_gateway/model"
	"github.com/google/uuid"
)

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userIDString := ctx.Value(model.ContextUser)
	if userIDString == nil {
		return uuid.Nil, ErrMissingAuthorizationToken
	}

	userID, err := uuid.Parse(fmt.Sprintf("%v", userIDString))
	if err != nil {
		return uuid.Nil, ErrInvalidAuthorizationToken
	}

	return userID, nil
}
