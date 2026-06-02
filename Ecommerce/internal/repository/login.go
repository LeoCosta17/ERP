package repository

import (
	"context"
	"database/sql"

	"github.com/leona/ecommerce/internal/model"
)

type LoginRepository struct {
	db *sql.DB
}

func (l *LoginRepository) Login(ctx context.Context, email, senha string) (*model.UsuarioLogin, error) {

	stmt, err := l.db.PrepareContext(ctx, "SELECT id, nome, email, senha FROM tb_usuarios WHERE email = ? and ativo = true")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var usuarioLogin model.UsuarioLogin

	err = stmt.QueryRowContext(ctx, email).Scan(&usuarioLogin.ID, &usuarioLogin.Nome, &usuarioLogin.Email, &usuarioLogin.Senha)
	if err != nil {
		return nil, err
	}

	return &usuarioLogin, nil
}
