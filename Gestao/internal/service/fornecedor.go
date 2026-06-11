package service

import (
	"context"
	"database/sql"
	"gestao/internal/model"
	"gestao/internal/repository"
	"gestao/pkg/dbhelper"
)

type FornecedorService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (s *FornecedorService) ListarFornecedores(ctx context.Context, busca string) ([]*model.Fornecedor, error) {
	var fornecedores []*model.Fornecedor
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		fornecedores, errTx = s.repository.Fornecedores.ListarFornecedores(ctx, tx, busca)
		return errTx
	})
	return fornecedores, err
}

// CriarFornecedor gerencia a regra de negócio para criar um fornecedor.
// Valida dados (se necessário) e abre a transação para o repositório.
func (s *FornecedorService) CriarFornecedor(ctx context.Context, f *model.Fornecedor) (*model.Fornecedor, error) {
	if err := f.Validar(); err != nil {
		return nil, err
	}

	var fornecedorCriado *model.Fornecedor
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		fornecedorCriado, errTx = s.repository.Fornecedores.CriarFornecedor(ctx, tx, f)
		return errTx
	})
	if err != nil {
		return nil, err
	}

	return fornecedorCriado, nil
}

func (s *FornecedorService) ObterFornecedorPorID(ctx context.Context, id int64) (*model.Fornecedor, error) {
	var fornecedor *model.Fornecedor
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		fornecedor, errTx = s.repository.Fornecedores.ObterFornecedorPorID(ctx, tx, id)
		return errTx
	})
	return fornecedor, err
}

func (s *FornecedorService) AtualizarFornecedor(ctx context.Context, id int64, f *model.Fornecedor) error {
	if err := f.Validar(); err != nil {
		return err
	}

	return dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		return s.repository.Fornecedores.AtualizarFornecedor(ctx, tx, id, f)
	})
}
