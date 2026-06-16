package controller

import (
	"gestaoadvocacia/internal/model"
	"gestaoadvocacia/internal/service"
	"gestaoadvocacia/pkg/requisicao"
	"gestaoadvocacia/pkg/resposta"
	"net/http"
)

type ClienteController struct {
	service *service.Service
}

func (c *ClienteController) CriarCliente(w http.ResponseWriter, r *http.Request) {
	var cliente model.Cliente

	if err := requisicao.ProcessarRequisicao(w, r, &cliente); err != nil {
		return
	}

	clienteCriado, err := c.service.Clientes.CriarCliente(r.Context(), &cliente)
	if err != nil {
		resposta.Padrao(w, http.StatusBadRequest, map[string]string{"erro": err.Error()})
		return
	}

	resposta.Padrao(w, http.StatusCreated, clienteCriado)
}

func (c *ClienteController) ListarClientes(w http.ResponseWriter, r *http.Request) {
	busca := r.URL.Query().Get("busca")

	clientes, err := c.service.Clientes.ListarClientes(r.Context(), busca)
	if err != nil {
		resposta.Padrao(w, http.StatusInternalServerError, map[string]string{"erro": "Erro ao buscar clientes: " + err.Error()})
		return
	}

	if clientes == nil {
		clientes = []model.Cliente{}
	}

	resposta.Padrao(w, http.StatusOK, clientes)
}

