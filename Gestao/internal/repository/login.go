package repository

import (
	"context"
	"database/sql"
)

type LoginRepository struct {
	db *sql.DB
}

func (r *LoginRepository) Login(ctx context.Context, tx *sql.Tx, email string) (uint64, uint64, string, string, error) {
	var senhaDB string
	var nome string
	var id uint64
	var id_empresa uint64
	err := tx.QueryRowContext(ctx, `
		select id, id_empresa, nome, senha from tb_usuarios_gestao
		where email = ?;
	`, email).Scan(&id, &id_empresa, &nome, &senhaDB)

	if err != nil {
		return 0, 0, "", "", err
	}

	return id, id_empresa, nome, senhaDB, nil
}
