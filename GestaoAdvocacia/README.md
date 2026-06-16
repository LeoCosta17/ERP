# ERP Gestão 🏢

Um Sistema de Gestão Empresarial (ERP) moderno, construído para ser rápido, seguro e escalável. Utiliza uma arquitetura híbrida de SPA (Single Page Application) onde o backend em Go gerencia as regras de negócio, banco de dados e rotas, enquanto o frontend utiliza Javascript puro para reatividade.

---

## 🚀 Tecnologias Utilizadas

### Backend (API & Renderização)
* **[Go (Golang)](https://go.dev/)**: Linguagem principal do servidor (v1.25.0).
* **[Chi Router](https://go-chi.io/)**: Roteador HTTP leve e rápido com suporte a Middlewares.
* **[MySQL](https://www.mysql.com/)**: Banco de dados relacional (via `go-sql-driver/mysql`).
* **Segurança**: Criptografia de senhas usando **Bcrypt** e geração de tokens de sessão com **JWT** (`golang-jwt/v5`).
* **Arquitetura**: Padrão limpo em camadas (`Controller` -> `Service` -> `Repository`).

### Frontend (Interface de Usuário)
* **HTML/Template do Go**: Usado para servir as "cascas" das páginas iniciais, permitindo o uso de Componentes (Herança de templates).
* **CSS Customizado**: Design premium com efeitos modernos como *Glassmorphism* e gradientes animados.
* **[Bootstrap 5](https://getbootstrap.com/) & Bootstrap Icons**: Utilizado para grid system de forma leve e componentes base.
* **Javascript (Vanilla)**: Consumo dinâmico das rotas da API (usando `fetch`) sem recarregar a página.

---

## 📁 Estrutura do Projeto

```text
Gestao/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada (Inicialização do servidor)
├── config/
│   └── env.go                   # Gerenciamento de Variáveis de Ambiente
├── db/
│   └── db.go                    # Conexão com o Banco de Dados MySQL
├── internal/
│   ├── auth/                    # Middlewares de Autenticação (Verificação de JWT)
│   ├── controller/              # Lida com requisições HTTP (API e Views)
│   ├── model/                   # Estruturas de dados (Structs) e validações
│   ├── repository/              # Queries de banco de dados (Acesso direto ao DB)
│   ├── router/                  # Definição das rotas e agrupamento de APIs
│   └── service/                 # Regras de negócio e orquestração
├── pkg/
│   ├── requisicao/              # Helpers para processar JSON de entrada
│   ├── resposta/                # Helpers para enviar JSON de saída
│   └── token/                   # Geração e validação de JWT
├── scripts/                     # Scripts SQL para criação de tabelas
└── web/
    ├── static/                  # Arquivos públicos
    │   └── css/                 # Folhas de estilo do sistema
    └── template/                # Arquivos HTML para renderização
        ├── components/          # Componentes reaproveitáveis (ex: Formulários, Sidebar)
        └── pages/               # Páginas base (ex: login.html, app.html)
```

---

## ⚙️ Como Executar o Projeto

1. **Configurar as Variáveis de Ambiente**:
   Certifique-se de que o arquivo `.envrc` (ou suas variáveis locais) está carregado com as configurações do banco de dados (`DB_USER`, `DB_PASS`, `DB_NAME`), porta do servidor (`API_PORT`) e a chave secreta JWT (`JWT_KEY`).

2. **Baixar Dependências**:
   No terminal, execute:
   ```bash
   go mod tidy
   ```

3. **Rodar o Servidor**:
   Sempre execute o projeto a partir da **raiz da pasta do projeto** (para que ele encontre as pastas `web/` e arquivos `.html` corretamente):
   ```bash
   go run ./cmd/api
   ```

4. **Acessar no Navegador**:
   * **Login**: [http://localhost:3000/](http://localhost:3000/) *(A porta pode variar conforme configurado no .envrc)*
   * **Dashboard**: [http://localhost:3000/dashboard](http://localhost:3000/dashboard)

---

## 🔐 Fluxo de Autenticação (Como funciona?)

1. As páginas visuais (Rotas Renderizadas em HTML) são **públicas**.
2. O usuário preenche o login e o Javascript faz um POST invisível para a rota de Funcionalidade (`/api/login`).
3. O Backend valida a senha em *Bcrypt*, gera um token *JWT* e devolve em formato JSON.
4. O Javascript salva esse token no `localStorage` do navegador do usuário.
5. Em navegações futuras, o Javascript injeta esse token no cabeçalho HTTP (`Authorization: Bearer <token>`) de forma silenciosa para capturar dados na API protegida.

---
*Desenvolvido em Junho de 2026.*
