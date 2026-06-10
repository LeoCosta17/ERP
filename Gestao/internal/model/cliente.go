package model

type TipoPessoa string
type IndContribuinte int

const (
	PessoaFisica   TipoPessoa = "PF"
	PessoaJuridica TipoPessoa = "PJ"

	ContribuinteICMS   IndContribuinte = 1
	ContribuinteIsento IndContribuinte = 2
	NaoContribuinte    IndContribuinte = 9
)

type Cliente struct {
	ID                int64             `json:"id,omitempty"`
	Nome              string            `json:"nome,omitempty"`
	Tipo              TipoPessoa        `json:"tipo,omitempty"`
	Email             string            `json:"email,omitempty"`
	Telefone          string            `json:"telefone,omitempty"`
	CPF               string            `json:"cpf,omitempty"`
	CNPJ              string            `json:"cnpj,omitempty"`
	Contribuinte      IndContribuinte   `json:"contribuinte,omitempty"`
	IsConsumidorFinal bool              `json:"is_consumidor_final,omitempty"`
	IE                string            `json:"ie,omitempty"`
	Enderecos         []EnderecoCliente `json:"endereco"`
	CreatedAt         string            `json:"created_at,omitempty"`
	UpdatedAt         string            `json:"updated_at,omitempty"`
}

type EnderecoCliente struct {
	ID              int64  `json:"id" db:"id"`
	IDCliente       int64  `json:"id_cliente"`
	CEP             string `json:"cep" db:"cep"`
	Logradouro      string `json:"logradouro" db:"logradouro"`
	Numero          string `json:"numero" db:"numero"`
	Bairro          string `json:"bairro" db:"bairro"`
	Municipio       string `json:"municipio" db:"municipio"`
	UF              string `json:"uf" db:"uf"`
	CodigoMunicipio string `json:"codigo_municipio" db:"codigo_municipio"`
	IsPrincipal     bool   `json:"is_principal" db:"is_principal"`
	CreatedAt       string `json:"created_at" db:"created_at"`
}
