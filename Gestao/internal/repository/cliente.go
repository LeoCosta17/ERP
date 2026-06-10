package repository

import (
	"context"
	"database/sql"
	"gestao/internal/model"
)

type ClienteRepository struct {
	db *sql.DB
}

func (r *ClienteRepository) CriarCliente(ctx context.Context, tx *sql.Tx, c *model.Cliente) (*model.Cliente, error) {
	var id int64

	query := `
		INSERT INTO tb_clientes (nome, tipo, email, telefone, cpf, cnpj, contribuinte, is_consumidor_final, ie)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, created_at;
	`
	err := tx.QueryRowContext(ctx, query, c.Nome, c.Tipo, c.Email, c.Telefone,
		c.CPF, c.CNPJ, c.Contribuinte, c.IsConsumidorFinal, c.IE).Scan(
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
			values (?, ?, ?, ?, ?, ?, ?, ?, ?)
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
