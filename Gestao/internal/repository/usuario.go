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
	var id uint64
	err := tx.QueryRowContext(ctx, `
		insert into tb_usuarios_gestao (id_empresa, nome, email, senha)
		values (?, ?, ?, ?)
		returning id;
	`, usuario.IDEmpresa, usuario.Nome, usuario.Email, usuario.Senha).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &model.UsuarioBasico{
		ID:        id,
		IDEmpresa: usuario.IDEmpresa,
		Nome:      usuario.Nome,
		Email:     usuario.Email,
	}, nil
}
