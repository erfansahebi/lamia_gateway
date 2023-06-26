package routes

import (
	"github.com/erfansahebi/lamia_gateway/edge/api/handlers"
	"github.com/erfansahebi/lamia_gateway/edge/api/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router, h handlers.Handler) {
	r.Post("/register", handlers.Wrap(h.Register))
	r.Post("/login", handlers.Wrap(h.Login))

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthenticateMiddleware(h.Di))

		r.Delete("/logout", handlers.Wrap(h.Logout))
	})

}
