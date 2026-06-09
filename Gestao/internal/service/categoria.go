package service

import (
	"context"
	"database/sql"
	"errors"
	"gestao/internal/model"
	"gestao/internal/repository"
)

type CategoriaService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (s *CategoriaService) CriarCategoria(ctx context.Context, c *model.CategoriaDebito) (*model.CategoriaDebito, error) {
	if c.Nome == "" {
		return nil, errors.New("o nome da categoria é obrigatório")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	categoriaCriada, err := s.repository.Categorias.CriarCategoria(ctx, tx, c)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return categoriaCriada, nil
}

func (s *CategoriaService) ListarCategorias(ctx context.Context) ([]*model.CategoriaDebito, error) {
	return s.repository.Categorias.ListarCategorias(ctx)
}
