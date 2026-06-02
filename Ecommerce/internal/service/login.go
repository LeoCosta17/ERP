package service

import (
	"context"
	"database/sql"

	"github.com/leona/ecommerce/internal/repository"
	"github.com/leona/ecommerce/internal/token"
)

type LoginService struct {
	repository *repository.Repository
	db         *sql.DB
}

func (l *LoginService) Login(ctx context.Context, email, senha string) (string, error) {

	usuario, err := l.repository.LoginRepository.Login(ctx, email, senha)
	if err != nil {
		return "", err
	}

	if err := usuario.CompararSenha(senha); err != nil {
		return "", err
	}

	token, err := token.GerarTokenJWT(usuario.ID, usuario.Nome)
	if err != nil {
		return "", err
	}

	return token, nil
}
