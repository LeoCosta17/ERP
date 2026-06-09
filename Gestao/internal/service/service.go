package service

import (
	"context"
	"database/sql"
	"gestao/internal/model"
	"gestao/internal/repository"
)

type Service struct {
	Usuarios interface {
		CriarUsuario(ctx context.Context, usuario *model.UsuarioCriar) (*model.UsuarioBasico, error)
	}
	Login interface {
		Login(ctx context.Context, usuario *model.UsuarioLogin) (uint64, string, error)
	}
	Fornecedores interface {
		CriarFornecedor(ctx context.Context, f *model.Fornecedor) (*model.Fornecedor, error)
		ListarFornecedores(ctx context.Context, busca string) ([]*model.Fornecedor, error)
	}
	Debitos interface {
		LancarDebito(ctx context.Context, debito *model.DebitoAvulsoCriar) error
		ListarDebitos(ctx context.Context, busca, vencimento, status string) ([]*model.Debito, error)
		PagarDebito(ctx context.Context, id int64) error
		EditarDebito(ctx context.Context, id int64, debito *model.DebitoAvulsoCriar) error
	}
	Categorias interface {
		CriarCategoria(ctx context.Context, c *model.CategoriaDebito) (*model.CategoriaDebito, error)
		ListarCategorias(ctx context.Context) ([]*model.CategoriaDebito, error)
	}
}

func NewService(repository *repository.Repository, db *sql.DB) *Service {
	return &Service{
		Usuarios: &UsuarioService{
			repository: repository,
			db:         db,
		},
		Login: &LoginService{
			repository: repository,
			db:         db,
		},
		Fornecedores: &FornecedorService{
			repository: repository,
			db:         db,
		},
		Categorias: &CategoriaService{
			repository: repository,
			db:         db,
		},
		Debitos: &DebitoService{
			repository: repository,
			db:         db,
		},
	}
}
