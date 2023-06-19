package common

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/edge/api/middleware"
	"github.com/google/uuid"
)

func GetUserIDFromContext(ctx context.Context) uuid.UUID {
	userData := ctx.Value(middleware.ContextUser)
	if userData == nil {
		return uuid.Nil
	}

	userID, err := uuid.Parse("userData")
	if err != nil {
		return uuid.Nil
	}

	return userID
}
