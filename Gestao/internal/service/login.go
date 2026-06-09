package service

import (
	"context"
	"database/sql"
	"errors"
	"gestao/internal/model"
	"gestao/internal/repository"
)

type LoginService struct {
	db         *sql.DB
	repository *repository.Repository
}

func (s *LoginService) Login(ctx context.Context, usuario *model.UsuarioLogin) (uint64, string, error) {
	if err := usuario.Validar(); err != nil {
		return 0, "", err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, "", err
	}
	defer tx.Rollback()

	id, nome, senhaDB, err := s.repository.Login.Login(ctx, tx, usuario.Email)
	if err != nil {
		return 0, "", err
	}

	if err := usuario.ValidarSenha(senhaDB); err != nil {
		return 0, "", errors.New("dados login inválidos")
	}

	if err := tx.Commit(); err != nil {
		return 0, "", err
	}

	return id, nome, nil
}
