package controller

import (
	"gestao/internal/model"
	"gestao/internal/service"
	"gestao/pkg/requisicao"
	"gestao/pkg/resposta"
	"net/http"
)

type UsuarioController struct {
	service *service.Service
}

func (c *UsuarioController) CriarUsuario(w http.ResponseWriter, r *http.Request) {

	var usuarioCriar model.UsuarioCriar
	if err := requisicao.ProcessarRequisicao(w, r, &usuarioCriar); err != nil {
		// ProcessarRequisicao já enviou um HTTP 400 caso tenha falhado
		return
	}

	usuarioCriado, err := c.service.Usuarios.CriarUsuario(r.Context(), &usuarioCriar)
	if err != nil {
		resposta.Padrao(w, http.StatusBadRequest, map[string]string{"erro": err.Error()})
		return
	}

	resposta.Padrao(w, http.StatusCreated, usuarioCriado)
}
