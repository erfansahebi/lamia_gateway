package routes

import (
	"github.com/erfansahebi/lamia_gateway/edge/api/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router, h handlers.Handler) {
	r.Post("/register", handlers.Wrap(h.Register))
	r.Post("/Login", handlers.Wrap(h.Login))
}
