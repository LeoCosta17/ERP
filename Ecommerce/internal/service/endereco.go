package service

import (
	"context"
	"database/sql"

	"github.com/leona/ecommerce/internal/model"
	"github.com/leona/ecommerce/internal/repository"
)

type EnderecoService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (s *EnderecoService) AdicionarEndereco(ctx context.Context, usuarioID int, cep, numero, complemento string) (*model.EnderecoUsuarioSimples, error) {

	endereco, err := ConsultaCEP(cep)
	if err != nil {
		return nil, err
	}

	endereco.Numero = numero
	endereco.Complemento = complemento

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	enderecoCriado, err := s.repository.EnderecoRepository.AdicionarEndereco(ctx, tx, usuarioID, endereco)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return enderecoCriado, nil

}
