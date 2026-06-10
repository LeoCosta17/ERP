package service

import (
	"context"
	"database/sql"
	"gestao/internal/model"
	"gestao/internal/repository"
)

type ClienteService struct {
	db         *sql.DB
	repository *repository.Repository
}

func (s *ClienteService) CriarCliente(ctx context.Context, c *model.Cliente) (*model.Cliente, error) {

	if err := c.Validar(); err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	cliente, err := s.repository.Clientes.CriarCliente(ctx, tx, c)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return cliente, nil
}
