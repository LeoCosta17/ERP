package repository

import (
	"context"
	"database/sql"
	"gestao/internal/model"
)

type UsuarioRepository struct {
	db *sql.DB
}

func (r *UsuarioRepository) CriarUsuario(ctx context.Context, tx *sql.Tx, usuario *model.UsuarioCriar) (*model.UsuarioBasico, error) {
	var id int64
	err := tx.QueryRowContext(ctx, `
		INSERT INTO tb_usuarios_gestao (nome, email, senha)
		VALUES ($1, $2, $3)
		RETURNING id;
	`, usuario.Nome, usuario.Email, usuario.Senha).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &model.UsuarioBasico{
		ID:    id,
		Nome:  usuario.Nome,
		Email: usuario.Email,
	}, nil
}

func (r *UsuarioRepository) BuscarUsuarioPorID(ctx context.Context, usuarioID int) (*model.Usuario, error) {
	var usuario model.Usuario
	err := r.db.QueryRowContext(ctx, `select id, nome, email, telefone from tb_usuarios_gestao where id = $1`).Scan(
		&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Telefone)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}
