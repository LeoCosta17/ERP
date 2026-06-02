package service

import (
	"context"
	"database/sql"

	"github.com/leona/ecommerce/internal/model"
	"github.com/leona/ecommerce/internal/repository"
)

type Service struct {
	UsuarioService interface {
		CriarUsuario(ctx context.Context, usuario *model.UsuarioCriar) (*model.UsuarioPublico, error)
		ValidarContaUsuario(ctx context.Context, token string) error
		AtualizarUsuario(ctx context.Context, usuario *model.UsuarioAtualizar, idUsuario int) (uint64, error)
	}
	LoginService interface {
		Login(ctx context.Context, email, senha string) (string, error)
	}
	EnderecoService interface {
		AdicionarEndereco(ctx context.Context, usuarioID int, cep, numero, complemento string) (*model.EnderecoUsuarioSimples, error)
	}
}

func NewService(repository *repository.Repository, db *sql.DB) *Service {
	return &Service{
		UsuarioService: &UsuarioService{repository: repository, db: db},
		LoginService:   &LoginService{repository: repository, db: db},
	}
}
