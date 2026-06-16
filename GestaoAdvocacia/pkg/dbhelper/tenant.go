package dbhelper

import (
	"context"
	"database/sql"
	"fmt"
)

type contextKey string

const SchemaKey = contextKey("schema_name")

// RunInTenantTx abre uma transação, configura o search_path para o schema do tenant 
// e executa a função informada. Dá commit em caso de sucesso ou rollback em erro.
func RunInTenantTx(ctx context.Context, db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Repassa o pânico após fazer rollback de segurança
		}
	}()

	// Pega o nome do schema do contexto (injetado pelo Middleware)
	schemaName, ok := ctx.Value(SchemaKey).(string)
	if !ok || schemaName == "" {
		schemaName = "public" // Fallback seguro
	}

	// Configura o search_path para isolar a transação
	query := fmt.Sprintf("SET search_path TO %s", schemaName)
	if _, err := tx.ExecContext(ctx, query); err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao configurar schema %s: %w", schemaName, err)
	}

	// Executa a função de negócio
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	// Commita se tudo deu certo
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("erro ao confirmar transação: %w", err)
	}

	return nil
}

