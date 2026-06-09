package service

import (
	"context"
	"database/sql"
	"gestao/internal/model"
	"gestao/internal/repository"
)

type FornecedorService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (s *FornecedorService) ListarFornecedores(ctx context.Context, busca string) ([]*model.Fornecedor, error) {
	return s.repository.Fornecedores.ListarFornecedores(ctx, busca)
}

// CriarFornecedor gerencia a regra de negócio para criar um fornecedor.
// Valida dados (se necessário) e abre a transação para o repositório.
func (s *FornecedorService) CriarFornecedor(ctx context.Context, f *model.Fornecedor) (*model.Fornecedor, error) {
	if err := f.Validar(); err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	fornecedorCriado, err := s.repository.Fornecedores.CriarFornecedor(ctx, tx, f)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return fornecedorCriado, nil
}
