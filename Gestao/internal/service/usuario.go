package service

import (
	"context"
	"database/sql"
	"gestao/internal/model"
	"gestao/internal/repository"
)

type UsuarioService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (s *UsuarioService) CriarUsuario(ctx context.Context, usuario *model.UsuarioCriar) (*model.UsuarioBasico, error) {

	if err := usuario.Validar(); err != nil {
		return nil, err
	}

	if err := usuario.HashSenha(); err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	usuarioCriado, err := s.repository.Usuarios.CriarUsuario(ctx, tx, usuario)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return usuarioCriado, nil
}
