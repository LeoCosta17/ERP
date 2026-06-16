package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

// IniciarTabelas lê todos os arquivos .sql da pasta scripts e os executa no banco de dados.
func IniciarTabelas(db *sql.DB) error {
	scriptsDir := "scripts"

	// Ordem correta para evitar erros de Chave Estrangeira (Foreign Key)
	arquivos := []string{
		"cria_tb_empresas.sql",
		"cria_tb_categorias_debito.sql",
		"cria_tb_clientes.sql",
		"cria_tb_fornecedores.sql",
		"cria_tb_usuarios_gestao.sql",
		"cria_tb_credenciais_empresa.sql",
		"cria_tb_dados_fiscais_empresa.sql",
		"cria_tb_endereco_empresa.sql",
		"cria_tb_enderecos_cliente.sql",
		"cria_tb_enderecos_fornecedor.sql",
		"cria_tb_telefones_fornecedor.sql",
		"cria_tb_debitos.sql",
	}

	fmt.Println("Iniciando verificação de tabelas no banco de dados...")

	for _, arquivo := range arquivos {
		path := filepath.Join(scriptsDir, arquivo)
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("erro ao ler arquivo %s: %w", path, err)
		}

		// Executa o script inteiro no banco
		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("erro ao executar script %s: %w", path, err)
		}
		fmt.Printf("-> Script %s executado com sucesso.\n", arquivo)
	}

	fmt.Println("Todas as tabelas foram iniciadas/verificadas com sucesso!")
	return nil
}
