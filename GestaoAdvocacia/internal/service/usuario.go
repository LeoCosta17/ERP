package service

import (
	"context"
	"database/sql"
	"gestaoadvocacia/internal/model"
	"gestaoadvocacia/internal/repository"
	"gestaoadvocacia/pkg/dbhelper"
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

	var usuarioCriado *model.UsuarioBasico
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		usuarioCriado, errTx = s.repository.Usuarios.CriarUsuario(ctx, tx, usuario)
		return errTx
	})
	if err != nil {
		return nil, err
	}

	return usuarioCriado, nil
}

