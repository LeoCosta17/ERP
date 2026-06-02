package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/leona/ecommerce/internal/model"
	"github.com/leona/ecommerce/internal/service"
	"github.com/leona/ecommerce/pkg/requisicao"
)

type UsuarioController struct {
	service *service.Service
}

// Recebe a requisição para criar um novo usuário, chama o serviço correspondente e retorna a resposta adequada
func (c *UsuarioController) CriarUsuario(w http.ResponseWriter, r *http.Request) {

	var usuario = &model.UsuarioCriar{}

	if err := requisicao.LerRequisicao(w, r, usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usuarioCriado, err := c.service.UsuarioService.CriarUsuario(r.Context(), usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarioCriado)
}

// Recebe a requisição para validar a conta do usuário, chama o serviço correspondente e retorna a resposta adequada
func (c *UsuarioController) ValidarContaUsuario(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("token")

	if err := c.service.UsuarioService.ValidarContaUsuario(r.Context(), token); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Conta validada com sucesso!"))
}

// Recebe a requisição para atualizar um usuário existente, chama o serviço correspondente e
// retorna a resposta adequada
func (c *UsuarioController) AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	usuarioRequisicaoID := uint64(ctx.Value("usuario_id").(float64))
	if usuarioRequisicaoID == 0 {
		http.Error(w, "ID do usuário não encontrado no contexto", http.StatusUnauthorized)
		return
	}

	usuarioID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID do usuário é inválido", http.StatusBadRequest)
		return
	}

	fmt.Printf("ID do usuário na requisição: %d\n", usuarioRequisicaoID)
	fmt.Printf("ID do usuário na URL: %s\n", r.PathValue("id"))

	if usuarioRequisicaoID != uint64(usuarioID) {
		http.Error(w, "Acesso negado: você só pode atualizar seu próprio perfil", http.StatusForbidden)
		return
	}

	var usuario = &model.UsuarioAtualizar{}

	if err := requisicao.LerRequisicao(w, r, usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rowsAffected, err := c.service.UsuarioService.AtualizarUsuario(ctx, usuario, usuarioID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"rows_affected": rowsAffected,
	})
}
