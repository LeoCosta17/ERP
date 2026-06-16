package service

import (
	"context"
	"database/sql"
	"gestaoadvocacia/internal/model"
	"gestaoadvocacia/internal/repository"
	"gestaoadvocacia/pkg/dbhelper"
)

type DebitoService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (s *DebitoService) LancarDebito(ctx context.Context, debito *model.DebitoAvulsoCriar) error {
	if err := debito.Validar(); err != nil {
		return err
	}

	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		return s.repository.Debitos.LancarDebito(ctx, tx, debito)
	})
	return err
}

func (s *DebitoService) ListarDebitos(ctx context.Context, busca, vencimento, status string) ([]*model.Debito, error) {
	var debitos []*model.Debito
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		debitos, errTx = s.repository.Debitos.ListarDebitos(ctx, tx, busca, vencimento, status)
		return errTx
	})
	return debitos, err
}

func (s *DebitoService) PagarDebito(ctx context.Context, id int64) error {
	return dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		return s.repository.Debitos.PagarDebito(ctx, tx, id)
	})
}

func (s *DebitoService) EditarDebito(ctx context.Context, id int64, debito *model.DebitoAvulsoCriar) error {
	if err := debito.Validar(); err != nil {
		return err
	}

	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		return s.repository.Debitos.EditarDebito(ctx, tx, id, debito)
	})
	return err
}

