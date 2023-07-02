package middleware

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/common"
	"github.com/erfansahebi/lamia_gateway/di"
	"github.com/erfansahebi/lamia_gateway/model"
	authProto "github.com/erfansahebi/lamia_shared/go/proto/auth"
	"net/http"
)

func AuthenticateMiddleware(container di.DIContainerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			authorizationToken := r.Header.Get("Authorization")

			if authorizationToken == "" {
				http.Error(rw, common.ErrMissingAuthorizationToken.Error(), http.StatusUnauthorized)
				return
			} else {
				authResponse, err := container.Services().Auth().Client().Authenticate(r.Context(), &authProto.AuthenticateRequest{
					AuthorizationToken: authorizationToken,
				})
				if err != nil {
					http.Error(rw, common.ErrUnAuthorized.Error(), http.StatusUnauthorized)
					return
				}

				ctx := context.WithValue(r.Context(), model.ContextUser, authResponse.Id)

				next.ServeHTTP(rw, r.WithContext(ctx))
			}
		})
	}
}
