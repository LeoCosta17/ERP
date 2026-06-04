package service

import (
	"context"
	"database/sql"
	"gestao/internal/model"
	"gestao/internal/repository"
)

type Service struct {
	Usuarios interface {
		CriarUsuario(ctx context.Context, usuario *model.UsuarioCriar) (*model.UsuarioBasico, error)
	}
	Login interface {
		Login(ctx context.Context, usuario *model.UsuarioLogin) (uint64, uint64, string, error)
	}
}

func NewService(repository *repository.Repository, db *sql.DB) *Service {
	return &Service{
		Usuarios: &UsuarioService{
			repository: repository,
			db:         db,
		},
		Login: &LoginService{
			repository: repository,
			db:         db,
		},
	}
}
