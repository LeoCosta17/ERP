package repository

import (
	"context"
	"database/sql"
	"errors"
	"gestao/internal/model"
)

type DebitoRepository struct {
	db *sql.DB
}

var (
	DEBITO_QUITADO        = errors.New("débito quitado")
	DEBITO_NAO_ENCONTRADO = errors.New("débito não encontrado")
)

func (r *DebitoRepository) LancarDebito(ctx context.Context, tx *sql.Tx, debito *model.DebitoAvulsoCriar) error {
	query := `
		INSERT INTO tb_debitos (
			id_fornecedor, id_categoria, descricao, 
			nr_documento, nr_nota_fiscal, valor, dt_entrada, dt_vencimento, 
			nr_parcela, nr_total_parcelas, status
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'PENDENTE')
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		debito.IDFornecedor, debito.IDCategoria, debito.Descricao,
		debito.NrDocumento, debito.NrNotaFiscal, debito.Valor, debito.DtEntrada, debito.DtVencimento,
		debito.NrParcela, debito.NrTotalParcelas,
	)
	return err
}

func (r *DebitoRepository) ListarDebitos(ctx context.Context, busca, vencimento, status string) ([]*model.Debito, error) {
	query := `
		SELECT d.id, d.id_fornecedor, d.id_categoria, d.descricao, d.nr_documento, d.nr_nota_fiscal, 
		       d.valor, d.dt_entrada, d.dt_vencimento, d.nr_parcela, d.nr_total_parcelas, d.status, d.created_at, d.updated_at,
		       f.id, f.razao_social, f.cnpj
		FROM tb_debitos d
		JOIN tb_fornecedores f ON d.id_fornecedor = f.id
		WHERE 1=1
	`
	var args []interface{}

	if busca != "" {
		query += " AND (f.razao_social LIKE ? OR f.id = ?)"
		buscaParam := "%" + busca + "%"
		args = append(args, buscaParam, busca)
	}

	if vencimento != "" {
		query += " AND d.dt_vencimento = ?"
		args = append(args, vencimento)
	}

	if status != "" {
		query += " AND d.status = ?"
		args = append(args, status)
	} else {
		query += " AND d.status = 'PENDENTE'"
	}

	query += " ORDER BY d.dt_vencimento ASC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var debitos []*model.Debito
	for rows.Next() {
		d := &model.Debito{}
		f := &model.Fornecedor{}

		err := rows.Scan(
			&d.ID, &d.IDFornecedor, &d.IDCategoria, &d.Descricao, &d.NrDocumento, &d.NrNotaFiscal,
			&d.Valor, &d.DtEntrada, &d.DtVencimento, &d.NrParcela, &d.NrTotalParcelas, &d.Status, &d.CreatedAt, &d.UpdatedAt,
			&f.ID, &f.RazaoSocial, &f.CNPJ,
		)
		if err != nil {
			return nil, err
		}

		d.Fornecedor = f
		debitos = append(debitos, d)
	}

	return debitos, nil
}

func (r *DebitoRepository) PagarDebito(ctx context.Context, tx *sql.Tx, id int64) error {

	query := `SELECT status from tb_debitos WHERE id = ?`
	stmt, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var status string
	if stmt.Next() {
		if err := stmt.Scan(&status); err != nil {
			return err
		}
	}

	if status != "PENDENTE" {
		return DEBITO_QUITADO
	}

	query = "UPDATE tb_debitos SET status = 'PAGO' WHERE id = ?"

	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return DEBITO_NAO_ENCONTRADO
	}
	return nil
}

func (r *DebitoRepository) EditarDebito(ctx context.Context, tx *sql.Tx, id int64, debito *model.DebitoAvulsoCriar) error {

	query := `SELECT status from tb_debitos WHERE id = ?`
	stmt, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var status string
	if stmt.Next() {
		if err := stmt.Scan(&status); err != nil {
			return err
		}
	}

	if status != "PENDENTE" {
		return DEBITO_QUITADO
	}

	query = `
		UPDATE tb_debitos SET 
			id_fornecedor = ?, id_categoria = ?, descricao = ?, 
			nr_documento = ?, nr_nota_fiscal = ?, valor = ?, dt_entrada = ?, dt_vencimento = ?, 
			nr_parcela = ?, nr_total_parcelas = ?
		WHERE id = ?
	`
	result, err := tx.ExecContext(ctx, query,
		debito.IDFornecedor, debito.IDCategoria, debito.Descricao,
		debito.NrDocumento, debito.NrNotaFiscal, debito.Valor, debito.DtEntrada, debito.DtVencimento,
		debito.NrParcela, debito.NrTotalParcelas, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return DEBITO_NAO_ENCONTRADO
	}
	return nil
}
