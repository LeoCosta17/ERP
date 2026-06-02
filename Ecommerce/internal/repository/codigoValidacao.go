package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type CodigoValidacaoRepository struct {
	db *sql.DB
}

// Cria um novo código de validação para o usuario, associando-o ao ID do usuario e ao tipo de validação (ex: validar conta, resetar senha, etc),
// e armazenando-o no banco de dados para posterior verificação e validação do usuario
func (r *CodigoValidacaoRepository) CriarCodigoValidacao(ctx context.Context, tx *sql.Tx, idUsuario int, codigo string, tipo string) error {
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO tb_codigos_validacao (usuario_id, codigo, tipo) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, idUsuario, codigo, tipo)
	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Printf("Código de validação criado com ID: %d\n", lastInsertId)
	return nil

}

// Valida o código de validação do usuario, verificando se ele existe no banco de dados,
// se é do tipo correto, se ainda não foi usado e se ainda não expirou, retornando o ID do usuario associado ao código para que sua conta possa ser ativada ou seu acesso possa ser liberado, dependendo do tipo de validação
func (r *CodigoValidacaoRepository) ValidarCodigo(ctx context.Context, tx *sql.Tx, token string) (*int, error) {

	stmt, err := tx.PrepareContext(ctx, "select usuario_id from tb_codigos_validacao where codigo = ? and tipo = ? and usado = false and expira_em > now()")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var idUsuario int
	err = stmt.QueryRowContext(ctx, token, "ATIVACAO").Scan(&idUsuario)
	if err != nil {
		return nil, err
	}

	stmt, err = tx.PrepareContext(ctx, "update tb_codigos_validacao set usado = true where codigo = ? and tipo = ? and usado = false and expira_em > now()")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, token, "ATIVACAO")
	if err != nil {
		return nil, err
	}

	return &idUsuario, nil
}
