package repository

import (
	"context"
	"database/sql"
	"gestao/internal/model"
)

type ClienteRepository struct {
	db *sql.DB
}

func nullIfEmpty(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

func nullIfZeroInt(i model.IndContribuinte) interface{} {
	if i == 0 {
		return nil
	}
	return i
}

func (r *ClienteRepository) CriarCliente(ctx context.Context, tx *sql.Tx, c *model.Cliente) (*model.Cliente, error) {
	var id int64

	query := `
		INSERT INTO tb_clientes (nome, tipo, email, telefone, cpf, cnpj, contribuinte, is_consumidor_final, ie)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at;
	`
	err := tx.QueryRowContext(ctx, query, c.Nome, c.Tipo, nullIfEmpty(c.Email), nullIfEmpty(c.Telefone),
		nullIfEmpty(c.CPF), nullIfEmpty(c.CNPJ), nullIfZeroInt(c.Contribuinte), c.IsConsumidorFinal, nullIfEmpty(c.IE)).Scan(
		&id, &c.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	c.ID = id

	for i, end := range c.Enderecos {
		var endID int64
		query = `
			insert into tb_enderecos_clientes(id_cliente, cep, logradouro, numero, bairro, municipio, uf, codigo_municipio, is_principal)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			returning id, created_at;
		`
		err := tx.QueryRowContext(
			ctx, query, c.ID, end.CEP, end.Logradouro,
			end.Numero, end.Bairro, end.Municipio, end.UF,
			end.CodigoMunicipio, end.IsPrincipal).Scan(
			&endID, &end.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		c.Enderecos[i].ID = endID
		c.Enderecos[i].IDCliente = c.ID
		c.Enderecos[i].CreatedAt = end.CreatedAt
	}

	return c, nil
}

func (r *ClienteRepository) ListarClientes(ctx context.Context, busca string) ([]model.Cliente, error) {
	query := `
		SELECT id, nome, tipo, email, telefone, cpf, cnpj 
		FROM tb_clientes
	`
	var rows *sql.Rows
	var err error

	if busca != "" {
		query += " WHERE nome ILIKE $1 OR cpf ILIKE $2 OR cnpj ILIKE $3"
		buscaParam := "%" + busca + "%"
		query += " ORDER BY id DESC"
		rows, err = r.db.QueryContext(ctx, query, buscaParam, buscaParam, buscaParam)
	} else {
		query += " ORDER BY id DESC"
		rows, err = r.db.QueryContext(ctx, query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []model.Cliente
	for rows.Next() {
		var c model.Cliente
		var email, telefone, cpf, cnpj sql.NullString
		
		if err := rows.Scan(
			&c.ID, &c.Nome, &c.Tipo, &email, &telefone, &cpf, &cnpj,
		); err != nil {
			return nil, err
		}
		
		c.Email = email.String
		c.Telefone = telefone.String
		c.CPF = cpf.String
		c.CNPJ = cnpj.String
		
		clientes = append(clientes, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clientes, nil
}
