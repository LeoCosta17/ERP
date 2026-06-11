package service

import (
	"context"
	"database/sql"
	"gestao/internal/model"
	"gestao/internal/repository"
	"gestao/pkg/dbhelper"
)

type ClienteService struct {
	db         *sql.DB
	repository *repository.Repository
}

func (s *ClienteService) CriarCliente(ctx context.Context, c *model.Cliente) (*model.Cliente, error) {

	if err := c.Validar(); err != nil {
		return nil, err
	}

	var clienteCriado *model.Cliente
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		clienteCriado, errTx = s.repository.Clientes.CriarCliente(ctx, tx, c)
		return errTx
	})
	if err != nil {
		return nil, err
	}

	return clienteCriado, nil
}
