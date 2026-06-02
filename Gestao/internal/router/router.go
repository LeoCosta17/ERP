package router

import (
	"gestao/config"
	"gestao/internal/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CarregarRotas(c *controller.Controller) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(config.GetInt("ROUTER_MAX_REQUESTS_PER_MINUTE", 10)))

	return r
}
