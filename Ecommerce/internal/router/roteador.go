package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	appMiddleware "github.com/leona/ecommerce/internal/middleware"

	"github.com/leona/ecommerce/configs"
	"github.com/leona/ecommerce/internal/controller"
)

func CarregarRotas(c *controller.Controller) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(configs.GetInt("ROUTER_MAX_REQUESTS_PER_MINUTE", 10)))

	r.Post("/", c.LoginController.Login)
	r.Post("/usuarios", c.UsuarioController.CriarUsuario)
	r.Get("/ativacao/usuario", c.UsuarioController.ValidarContaUsuario)
	r.Put("/usuarios/{id}", appMiddleware.Autenticar(c.UsuarioController.AtualizarUsuario))

	return r
}
