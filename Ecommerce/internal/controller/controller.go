package controller

import (
	"net/http"

	"github.com/leona/ecommerce/internal/service"
)

type Controller struct {
	LoginController interface {
		Login(w http.ResponseWriter, r *http.Request)
	}
	UsuarioController interface {
		CriarUsuario(w http.ResponseWriter, r *http.Request)
		ValidarContaUsuario(w http.ResponseWriter, r *http.Request)
		AtualizarUsuario(w http.ResponseWriter, r *http.Request)
	}
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		LoginController:   &LoginController{service: service},
		UsuarioController: &UsuarioController{service: service},
	}
}
