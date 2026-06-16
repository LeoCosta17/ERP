package middleware

import (
	"context"
	"gestao/pkg/dbhelper"
	"net/http"
	"strings"
)

// TenantMiddleware extrai o subdomínio da requisição e injeta o schema_name no contexto.
func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := r.Host // ex: "tobias.gestao.com" ou "localhost:9000"

		var schemaName string

		// Se estiver rodando localmente, vai para o schema public (painel/testes)
		if strings.HasPrefix(host, "localhost") || strings.HasPrefix(host, "127.0.0.1") {
			schemaName = "public"
		} else {
			// Extrai o subdomínio (primeira parte antes do ponto)
			partes := strings.Split(host, ".")

			// Se tiver pelo menos "subdominio.dominio.com" (3 partes)
			// Ou "subdominio.dominio.com.br" (4 partes)
			if len(partes) > 2 {
				subdominio := partes[0]

				if subdominio == "admin" || subdominio == "www" {
					schemaName = "public"
				} else {
					schemaName = "schema_" + subdominio
				}
			} else {
				// Acesso direto sem subdomínio
				schemaName = "public"
			}
		}

		// Injeta no contexto
		ctx := context.WithValue(r.Context(), dbhelper.SchemaKey, schemaName)

		// Segue o fluxo passando o novo contexto
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
