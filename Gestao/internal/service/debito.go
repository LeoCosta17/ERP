package service

import (
	"context"
	"database/sql"
	"gestao/internal/model"
	"gestao/internal/repository"
)

type DebitoService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (s *DebitoService) LancarDebito(ctx context.Context, debito *model.DebitoAvulsoCriar) error {
	if err := debito.Validar(); err != nil {
		return err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := s.repository.Debitos.LancarDebito(ctx, tx, debito); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *DebitoService) ListarDebitos(ctx context.Context, busca, vencimento, status string) ([]*model.Debito, error) {
	return s.repository.Debitos.ListarDebitos(ctx, busca, vencimento, status)
}

func (s *DebitoService) PagarDebito(ctx context.Context, id int64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := s.repository.Debitos.PagarDebito(ctx, tx, id); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *DebitoService) EditarDebito(ctx context.Context, id int64, debito *model.DebitoAvulsoCriar) error {
	if err := debito.Validar(); err != nil {
		return err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := s.repository.Debitos.EditarDebito(ctx, tx, id, debito); err != nil {
		return err
	}

	return tx.Commit()
}
