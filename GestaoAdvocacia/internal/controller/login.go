package controller

import (
	"gestaoadvocacia/internal/model"
	"gestaoadvocacia/internal/service"
	"gestaoadvocacia/pkg/requisicao"
	"gestaoadvocacia/pkg/resposta"
	"gestaoadvocacia/pkg/token"
	"net/http"
)

type LoginController struct {
	service *service.Service
}

func (c *LoginController) Login(w http.ResponseWriter, r *http.Request) {

	var usuarioRequest model.UsuarioLogin

	if err := requisicao.ProcessarRequisicao(w, r, &usuarioRequest); err != nil {
		resposta.Padrao(w, http.StatusBadRequest, "erro ao ler json")
		return
	}

	id, nome, err := c.service.Login.Login(r.Context(), &usuarioRequest)
	if err != nil {
		resposta.Padrao(w, http.StatusUnauthorized, "dados login inválidos")
		return
	}

	tokenString, err := token.GerarTokenJWT(int(id), nome)
	if err != nil {
		resposta.Padrao(w, http.StatusInternalServerError, "erro ao gerar token")
		return
	}

	token := map[string]string{
		"token": tokenString,
	}

	resposta.Padrao(w, http.StatusOK, token)

}

