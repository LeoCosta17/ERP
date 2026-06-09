package repository

import (
	"context"
	"database/sql"
	"gestao/internal/model"
)

type FornecedorRepository struct {
	db *sql.DB
}

// NovoFornecedorRepository cria uma nova instância do repositório
func NovoFornecedorRepository(db *sql.DB) *FornecedorRepository {
	return &FornecedorRepository{
		db: db,
	}
}

func (r *FornecedorRepository) ListarFornecedores(ctx context.Context, busca string) ([]*model.Fornecedor, error) {
	query := `
		SELECT id, razao_social, cnpj, inscricao_estadual, email, created_at, updated_at
		FROM tb_fornecedores
	`
	var rows *sql.Rows
	var err error

	if busca != "" {
		query += " WHERE razao_social LIKE ? OR cnpj LIKE ?"
		buscaParam := "%" + busca + "%"
		rows, err = r.db.QueryContext(ctx, query, buscaParam, buscaParam)
	} else {
		rows, err = r.db.QueryContext(ctx, query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fornecedores []*model.Fornecedor
	for rows.Next() {
		f := &model.Fornecedor{}
		err := rows.Scan(&f.ID, &f.RazaoSocial, &f.CNPJ, &f.InscricaoEstadual, &f.Email, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			return nil, err
		}
		fornecedores = append(fornecedores, f)
	}
	return fornecedores, nil
}

// CriarFornecedor insere um novo fornecedor e seus relacionamentos (se existirem) no banco de dados.
// Recebe um tx *sql.Tx pois a inserção em múltiplas tabelas (fornecedores, endereços, telefones) deve ser transacional.
func (r *FornecedorRepository) CriarFornecedor(ctx context.Context, tx *sql.Tx, f *model.Fornecedor) (*model.Fornecedor, error) {
	var id int64

	// 1. Inserir Fornecedor
	queryFornecedor := `
		INSERT INTO tb_fornecedores (razao_social, cnpj, inscricao_estadual, email)
		VALUES (?, ?, ?, ?)
		RETURNING id, created_at, updated_at;
	`
	err := tx.QueryRowContext(ctx, queryFornecedor, f.RazaoSocial, f.CNPJ, f.InscricaoEstadual, f.Email).Scan(&id, &f.CreatedAt, &f.UpdatedAt)
	if err != nil {
		return nil, err
	}
	f.ID = id

	// 2. Inserir Endereços (Opcional)
	for i, end := range f.Enderecos {
		var endID int64
		queryEndereco := `
			INSERT INTO tb_enderecos_fornecedores (id_fornecedor, cep, logradouro, numero, bairro, municipio, uf, codigo_municipio, is_principal)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
			RETURNING id, created_at;
		`
		err = tx.QueryRowContext(ctx, queryEndereco,
			f.ID, end.CEP, end.Logradouro, end.Numero, end.Bairro, end.Municipio, end.UF, end.CodigoMunicipio, end.IsPrincipal,
		).Scan(&endID, &end.CreatedAt)
		
		if err != nil {
			return nil, err
		}
		
		// Atualiza os dados do item no slice
		f.Enderecos[i].ID = endID
		f.Enderecos[i].IDFornecedor = f.ID
		f.Enderecos[i].CreatedAt = end.CreatedAt
	}

	// 3. Inserir Telefones (Opcional)
	for i, tel := range f.Telefones {
		var telID int64
		queryTelefone := `
			INSERT INTO tb_telefones_fornecedores (id_fornecedor, ddd, numero)
			VALUES (?, ?, ?)
			RETURNING id, created_at;
		`
		err = tx.QueryRowContext(ctx, queryTelefone, f.ID, tel.DDD, tel.Numero).Scan(&telID, &tel.CreatedAt)
		if err != nil {
			return nil, err
		}
		
		// Atualiza os dados do item no slice
		f.Telefones[i].ID = telID
		f.Telefones[i].IDFornecedor = f.ID
		f.Telefones[i].CreatedAt = tel.CreatedAt
	}

	return f, nil
}
