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
	r.Get("/debitos", c.View.RenderizarDebitosPage)
	r.Get("/fornecedores", c.View.RenderizarFornecedoresPage)
	r.Get("/categorias-debito", c.View.RenderizarCategoriasPage)

	// rotas funcionalidades
	r.Route("/api", func(r chi.Router) {
		r.Post("/login", c.Login.Login)
		r.Post("/usuarios", c.Usuarios.CriarUsuario)
		r.Get("/fornecedores", auth.Autenticar(c.Fornecedores.ListarFornecedores))
		r.Post("/fornecedores", auth.Autenticar(c.Fornecedores.CriarFornecedor))
		r.Get("/categorias", auth.Autenticar(c.Categorias.ListarCategorias))
		r.Post("/categorias", auth.Autenticar(c.Categorias.CriarCategoria))
		r.Get("/debitos", auth.Autenticar(c.Debitos.ListarDebitos))
		r.Post("/debitos", auth.Autenticar(c.Debitos.CriarDebitoAvulso))
		r.Put("/debitos/{id}", auth.Autenticar(c.Debitos.EditarDebito))
		r.Put("/debitos/{id}/pagar", auth.Autenticar(c.Debitos.PagarDebito))
		r.Get("/dashboard/resumo", auth.Autenticar(c.Dashboard.ResumoDashboard))
	})

	return r
}
