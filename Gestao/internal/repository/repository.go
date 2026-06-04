package repository

import (
	"context"
	"database/sql"
	"gestao/internal/model"
)

type Repository struct {
	Login interface {
		Login(ctx context.Context, tx *sql.Tx, email string) (uint64, uint64, string, string, error)
	}
	Usuarios interface {
		CriarUsuario(ctx context.Context, tx *sql.Tx, usuario *model.UsuarioCriar) (*model.UsuarioBasico, error)
	}
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Usuarios: &UsuarioRepository{
			db: db,
		},
	}
}
