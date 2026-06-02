package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/leona/ecommerce/internal/model"
	"github.com/leona/ecommerce/internal/repository"
)

type UsuarioService struct {
	repository *repository.Repository
	db         *sql.DB
}

// Cria um novo usuario no banco de dados, gera seu código de validação e o envia ao usuario por email ou whatsapp, utilizando uma transação para garantir a integridade dos dados
func (s *UsuarioService) CriarUsuario(ctx context.Context, usuario *model.UsuarioCriar) (*model.UsuarioPublico, error) {

	if err := usuario.Validar(); err != nil {
		return nil, err
	}

	if err := usuario.HashSenha(); err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	usuarioCriado, err := s.repository.UsuarioRepository.CriarUsuario(ctx, tx, usuario)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tokenValidacao := uuid.New().String()

	if err := s.repository.CodigoValidacaoRepository.CriarCodigoValidacao(ctx, tx, usuarioCriado.ID, tokenValidacao, "ATIVACAO"); err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	if err := EnviarEmail(usuarioCriado.Nome, tokenValidacao, usuarioCriado.Email); err != nil {
		return nil, err
	}

	return usuarioCriado, nil
}

// Valida o código de validação do usuario, ativando sua conta no sistema e permitindo que ele faça login e utilize os serviços da plataforma
func (s *UsuarioService) ValidarContaUsuario(ctx context.Context, token string) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	idUsuario, err := s.repository.CodigoValidacaoRepository.ValidarCodigo(ctx, tx, token)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = s.repository.UsuarioRepository.ValidarContaUsuario(ctx, tx, *idUsuario)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Atualiza os dados do usuario, utilizando uma transação para garantir a integridade dos dados e validando os campos obrigatórios antes de realizar a atualização no banco de dados
func (s *UsuarioService) AtualizarUsuario(ctx context.Context, usuario *model.UsuarioAtualizar, idUsuario int) (uint64, error) {
	if err := usuario.Validar(); err != nil {
		return 0, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := s.repository.UsuarioRepository.AtualizarUsuario(ctx, tx, usuario, idUsuario)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
