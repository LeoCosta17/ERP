package router

import (
	"gestao/config"
	"gestao/internal/auth"
	"gestao/internal/controller"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CarregarRotas(c *controller.Controller) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(config.GetInt("ROUTER_MAX_REQUESTS_PER_MINUTE", 10)))

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filesDir := http.Dir(filepath.Join(workDir, "web", "static"))
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))

	// rotas renderizar páginas
	r.Get("/", c.View.RenderizarLoginPage)
	r.Get("/dashboard", c.View.RenderizarDashboardPage)

	// rotas funcionalidades
	r.Route("/api", func(r chi.Router) {
		r.Post("/login", c.Login.Login)
		r.Post("/usuarios", auth.Autenticar(c.Usuarios.CriarUsuario))
	})

	return r
}
