package api

import (
	"context"
	"fmt"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_gateway/di"
	"github.com/erfansahebi/lamia_gateway/edge/api/handlers"
	"github.com/erfansahebi/lamia_gateway/edge/api/routes"
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

		r.Route("/auth", func(r chi.Router) {

			routes.RegisterAuthRoutes(r, h)

		})

	})

	httpConfig := configuration.Server.HTTP
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", httpConfig.Host, httpConfig.Port), r); err != nil {
		os.Exit(1)
	}

}
