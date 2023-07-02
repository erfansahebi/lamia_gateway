package api

import (
	"context"
	"fmt"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_gateway/di"
	"github.com/erfansahebi/lamia_gateway/edge/api/handlers"
	"github.com/erfansahebi/lamia_gateway/edge/api/middleware"
	"github.com/erfansahebi/lamia_gateway/edge/api/routes"
	"github.com/erfansahebi/lamia_shared/go/log"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

func StartServer(ctx context.Context, configuration *config.Config) {
	h := handlers.Handler{
		AppCtx: ctx,
		Di:     di.NewDIContainer(ctx, configuration),
	}

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			r.Route("/auth", func(r chi.Router) {

				routes.RegisterAuthRoutes(r, h)

			})

			r.Group(func(r chi.Router) {

				r.Use(middleware.AuthenticateMiddleware(h.Di))

				r.Route("/user", func(r chi.Router) {

					routes.RegisterUserRoutes(r, h)

				})

				r.Route("/shop", func(r chi.Router) {

					routes.RegisterShopRoutes(r, h)

				})

			})

		})

	})

	httpConfig := configuration.Server.HTTP
	log.Infof(context.Background(), "API Server starting on: %s:%s", httpConfig.Host, httpConfig.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", httpConfig.Host, httpConfig.Port), r); err != nil {
		log.WithError(err).Errorf(context.Background(), "ListenAndServe failed")
		os.Exit(1)
	}

}
