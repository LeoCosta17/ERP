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

	var endID int64
	query = `
			insert into tb_enderecos_clientes(id_cliente, cep, logradouro, numero, bairro, municipio, uf, codigo_municipio, is_principal)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			returning id, created_at;
		`
	err = tx.QueryRowContext(
		ctx, query, c.ID, c.Endereco.CEP, c.Endereco.Logradouro,
		c.Endereco.Numero, c.Endereco.Bairro, c.Endereco.Municipio, c.Endereco.UF,
		c.Endereco.CodigoMunicipio, c.Endereco.IsPrincipal).Scan(
		&endID, &c.Endereco.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	c.Endereco.ID = endID
	c.Endereco.IDCliente = c.ID

	return c, nil
}

func (r *ClienteRepository) ListarClientes(ctx context.Context, tx *sql.Tx, busca string) ([]model.Cliente, error) {
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
		rows, err = tx.QueryContext(ctx, query, buscaParam, buscaParam, buscaParam)
	} else {
		query += " ORDER BY id DESC"
		rows, err = tx.QueryContext(ctx, query)
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

func (r *ClienteRepository) ObterClientePorID(ctx context.Context, tx *sql.Tx, id int64) (*model.Cliente, error) {
	query := `
		SELECT id, nome, tipo, email, telefone, cpf, cnpj, contribuinte, is_consumidor_final, ie, created_at, updated_at
		FROM tb_clientes
		WHERE id = $1
	`
	c := &model.Cliente{}
	var contribuinte sql.NullInt64
	var email, telefone, cpf, cnpj, ie sql.NullString

	err := tx.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.Nome, &c.Tipo, &email, &telefone, &cpf, &cnpj, &contribuinte, &c.IsConsumidorFinal, &ie, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	c.Email = email.String
	c.Telefone = telefone.String
	c.CPF = cpf.String
	c.CNPJ = cnpj.String
	c.IE = ie.String
	if contribuinte.Valid {
		c.Contribuinte = model.IndContribuinte(contribuinte.Int64)
	}

	// Buscar endereços
	queryEnderecos := `
		SELECT id, id_cliente, cep, logradouro, numero, bairro, municipio, uf, codigo_municipio, is_principal, created_at
		FROM tb_enderecos_clientes
		WHERE id_cliente = $1
	`
	err = tx.QueryRowContext(ctx, queryEnderecos, id).Scan(
		&c.Endereco.ID, &c.Endereco.IDCliente, &c.Endereco.CEP, &c.Endereco.Logradouro, &c.Endereco.Numero, &c.Endereco.Bairro,
		&c.Endereco.Municipio, &c.Endereco.UF, &c.Endereco.CodigoMunicipio, &c.Endereco.IsPrincipal, &c.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *ClienteRepository) AtualizarCliente(ctx context.Context, tx *sql.Tx, id int64, c *model.Cliente) error {
	query := `
		UPDATE tb_clientes
		SET nome = $1, tipo = $2, email = $3, telefone = $4, cpf = $5, cnpj = $6, contribuinte = $7, is_consumidor_final = $8, ie = $9, updated_at = CURRENT_TIMESTAMP
		WHERE id = $10
	`
	_, err := tx.ExecContext(ctx, query, c.Nome, c.Tipo, nullIfEmpty(c.Email), nullIfEmpty(c.Telefone),
		nullIfEmpty(c.CPF), nullIfEmpty(c.CNPJ), nullIfZeroInt(c.Contribuinte), c.IsConsumidorFinal, nullIfEmpty(c.IE), id)
	if err != nil {
		return err
	}

	// Deleta endereços atuais
	_, err = tx.ExecContext(ctx, "DELETE FROM tb_enderecos_clientes WHERE id_cliente = $1", id)
	if err != nil {
		return err
	}

	// Insere novos endereços
	queryEndereco := `
			insert into tb_enderecos_clientes(id_cliente, cep, logradouro, numero, bairro, municipio, uf, codigo_municipio, is_principal)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`
	_, err = tx.ExecContext(ctx, queryEndereco,
		id, c.Endereco.CEP, c.Endereco.Logradouro, c.Endereco.Numero, c.Endereco.Bairro,
		c.Endereco.Municipio, c.Endereco.UF, c.Endereco.CodigoMunicipio, c.Endereco.IsPrincipal,
	)
	if err != nil {
		return err
	}

	return nil
}
