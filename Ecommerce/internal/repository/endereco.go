package repository

import (
	"context"
	"database/sql"

	"github.com/leona/ecommerce/internal/model"
)

type EnderecoRepository struct {
	db *sql.DB
}

// Adiciona um endereço ao perfil do usuário, utilizando os dados retornados pela consulta
// ao CEP e o ID do usuário
func (r *EnderecoRepository) AdicionarEndereco(ctx context.Context, tx *sql.Tx, userID int, endereco *model.EnderecoConsultaCEP) (*model.EnderecoUsuarioSimples, error) {
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO tb_enderecos (usuario_id, cep, logradouro, numero, complemento, bairro, cidade, uf, ibge, gia, ddd, siafi) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, userID, endereco.CEP, endereco.Logradouro, endereco.Numero, endereco.Complemento, endereco.Bairro, endereco.Cidade, endereco.UF, endereco.Ibge, endereco.Gia, endereco.Ddd, endereco.Siafi)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	enderecoCriado := &model.EnderecoUsuarioSimples{
		ID:          int(id),
		IDusuario:   userID,
		CEP:         endereco.CEP,
		Logradouro:  endereco.Logradouro,
		Numero:      endereco.Numero,
		Complemento: endereco.Complemento,
		Bairro:      endereco.Bairro,
		Cidade:      endereco.Cidade,
	}
	return enderecoCriado, nil
}
