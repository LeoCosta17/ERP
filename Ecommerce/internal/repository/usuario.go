package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/leona/ecommerce/internal/model"
)

type UsuarioRepository struct {
	db *sql.DB
}

// Cria um novo usuario no banco de dados
func (r *UsuarioRepository) CriarUsuario(ctx context.Context, tx *sql.Tx, usuario *model.UsuarioCriar) (*model.UsuarioPublico, error) {
	// Implementar a lógica para criar um usuário no banco de dados

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO tb_usuarios (nome, cpf, email, senha) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, usuario.Nome, usuario.CPF, usuario.Email, usuario.Senha)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	usuarioCriado := &model.UsuarioPublico{
		ID:    int(lastInsertID),
		Nome:  usuario.Nome,
		Email: usuario.Email,
	}

	return usuarioCriado, nil
}

// Valida o código de validação do usuario, ativando sua conta no sistema e
// permitindo que ele faça login e utilize os serviços da plataforma
func (r *UsuarioRepository) ValidarContaUsuario(ctx context.Context, tx *sql.Tx, idUsuario int) error {

	stmt, err := tx.PrepareContext(ctx, "UPDATE tb_usuarios SET ativo = true WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, idUsuario)
	if err != nil {
		return err
	}

	return nil
}

// Atualiza os dados do usuario no banco de dados
func (r *UsuarioRepository) AtualizarUsuario(ctx context.Context, tx *sql.Tx, usuario *model.UsuarioAtualizar, idUsuario int) (uint64, error) {
	stmt, err := tx.PrepareContext(ctx, "UPDATE tb_usuarios SET nome = ?, cpf = ?, telefone = ?, email = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, usuario.Nome, usuario.CPF, usuario.Telefone, usuario.Email, idUsuario)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, errors.New("usuário não encontrado")
	}

	return uint64(rowsAffected), nil
}
