package routes

import (
	"github.com/erfansahebi/lamia_gateway/edge/api/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, h handlers.Handler) {
	r.Get("/", handlers.Wrap(h.UserDetail))
}
