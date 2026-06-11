package service

import (
	"context"
	"database/sql"
	"errors"
	"gestao/internal/model"
	"gestao/internal/repository"
	"gestao/pkg/dbhelper"
)

type LoginService struct {
	db         *sql.DB
	repository *repository.Repository
}

func (s *LoginService) Login(ctx context.Context, usuario *model.UsuarioLogin) (uint64, string, error) {
	if err := usuario.Validar(); err != nil {
		return 0, "", err
	}

	var id uint64
	var nome string
	var senhaDB string
	
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		id, nome, senhaDB, errTx = s.repository.Login.Login(ctx, tx, usuario.Email)
		return errTx
	})
	if err != nil {
		return 0, "", err
	}

	if err := usuario.ValidarSenha(senhaDB); err != nil {
		return 0, "", errors.New("dados login inválidos")
	}

	return id, nome, nil
}
