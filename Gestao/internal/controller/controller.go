package controller

import (
	"gestao/internal/service"
	"net/http"
)

type Controller struct {
	View interface {
		RenderizarLoginPage(w http.ResponseWriter, r *http.Request)
		RenderizarDashboardPage(w http.ResponseWriter, r *http.Request)
	}
	Login interface {
		Login(w http.ResponseWriter, r *http.Request)
	}
	Usuarios interface {
		CriarUsuario(w http.ResponseWriter, r *http.Request)
	}
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		View:     &ViewController{},
		Login:    &LoginController{service: service},
		Usuarios: &UsuarioController{service: service},
	}
}
