package controller

import (
	"net/http"

	"github.com/leona/ecommerce/internal/service"
	"github.com/leona/ecommerce/pkg/requisicao"
)

type LoginController struct {
	service *service.Service
}

type Payload struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {

	var payload Payload

	if err := requisicao.LerRequisicao(w, r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := l.service.LoginService.Login(r.Context(), payload.Email, payload.Senha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token": "` + token + `"}`))
}
