package service

import (
	"context"
	"database/sql"
	"errors"
	"gestaoadvocacia/internal/model"
	"gestaoadvocacia/internal/repository"
	"gestaoadvocacia/pkg/dbhelper"
)

type CategoriaService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (s *CategoriaService) CriarCategoria(ctx context.Context, c *model.CategoriaDebito) (*model.CategoriaDebito, error) {
	if c.Nome == "" {
		return nil, errors.New("o nome da categoria é obrigatório")
	}

	var categoriaCriada *model.CategoriaDebito
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		categoriaCriada, errTx = s.repository.Categorias.CriarCategoria(ctx, tx, c)
		return errTx
	})
	if err != nil {
		return nil, err
	}

	return categoriaCriada, nil
}

func (s *CategoriaService) ListarCategorias(ctx context.Context) ([]*model.CategoriaDebito, error) {
	var categorias []*model.CategoriaDebito
	err := dbhelper.RunInTenantTx(ctx, s.db, func(tx *sql.Tx) error {
		var errTx error
		categorias, errTx = s.repository.Categorias.ListarCategorias(ctx, tx)
		return errTx
	})
	return categorias, err
}

