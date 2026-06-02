package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/leona/ecommerce/internal/service"
	"github.com/leona/ecommerce/pkg/requisicao"
)

type EnderecoController struct {
	service *service.Service
}

// Recebe a requisição para adicionar um endereço ao perfil do usuário, chama o serviço correspondente e retorna a resposta adequada
func (c *UsuarioController) AdicionarEndereco(w http.ResponseWriter, r *http.Request) {

	type payload struct {
		CEP         string `json:"cep"`
		numero      string `json:"numero"`
		complemento string `json:"complemento"`
	}

	ctx := r.Context()

	var p payload

	if err := requisicao.LerRequisicao(w, r, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usuarioRequisicaoID := uint64(ctx.Value("usuario_id").(float64))
	usuarioID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID do usuário inválido", http.StatusBadRequest)
		return
	}

	if usuarioRequisicaoID != uint64(usuarioID) {
		http.Error(w, "Acesso negado: você não tem permissão para modificar este usuário", http.StatusForbidden)
		return
	}

	enderecoCriado, err := c.service.EnderecoService.AdicionarEndereco(ctx, usuarioID, p.CEP, p.numero, p.complemento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enderecoCriado)

}
