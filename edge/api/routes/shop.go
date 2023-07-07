package routes

import (
	"github.com/erfansahebi/lamia_gateway/edge/api/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterShopRoutes(r chi.Router, h handlers.Handler) {
	r.Post("/", handlers.Wrap(h.CreateShop))
	r.Get("/{shop_id}", handlers.Wrap(h.GetShop))
}
