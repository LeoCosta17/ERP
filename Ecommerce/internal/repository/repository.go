package repository

import (
	"context"
	"database/sql"

	"github.com/leona/ecommerce/internal/model"
)

type Repository struct {
	UsuarioRepository interface {
		CriarUsuario(ctx context.Context, tx *sql.Tx, usuario *model.UsuarioCriar) (*model.UsuarioPublico, error)
		ValidarContaUsuario(ctx context.Context, tx *sql.Tx, idUsuario int) error
		AtualizarUsuario(ctx context.Context, tx *sql.Tx, usuario *model.UsuarioAtualizar, idUsuario int) (uint64, error)
	}
	CodigoValidacaoRepository interface {
		CriarCodigoValidacao(ctx context.Context, tx *sql.Tx, idUsuario int, codigo string, tipo string) error
		ValidarCodigo(ctx context.Context, tx *sql.Tx, token string) (*int, error)
	}
	LoginRepository interface {
		Login(ctx context.Context, email, senha string) (*model.UsuarioLogin, error)
	}
	EnderecoRepository interface {
		AdicionarEndereco(ctx context.Context, tx *sql.Tx, userID int, endereco *model.EnderecoConsultaCEP) (*model.EnderecoUsuarioSimples, error)
	}
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UsuarioRepository:         &UsuarioRepository{db: db},
		CodigoValidacaoRepository: &CodigoValidacaoRepository{db: db},
		LoginRepository:           &LoginRepository{db: db},
	}
}
