package service

import (
	"database/sql"
	"gestao/internal/repository"
)

type Service struct {
}

func NewService(repository *repository.Repository, db *sql.DB) *Service {
	return &Service{}
}
