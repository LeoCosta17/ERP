package model

// Endereço retornado pela consulta ao CEP, sem os campos adicionais do cadastro do usuário
type EnderecoConsultaCEP struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	UF          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// Endereço que aparece no cadastro do usuário, com os campos adicionais da consulta ao CEP
// Será utilizado para uso em casos fiscais
type EnderecoUsuarioCompleto struct {
	IDusuario   int    `json:"id_usuario"`
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	UF          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Principal   bool   `json:"principal"`
}

// Endereço que aparece no cadastro do usuário, sem os campos adicionais da consulta ao CEP
type EnderecoUsuarioSimples struct {
	ID          int    `json:"id"`
	IDusuario   int    `json:"id_usuario"`
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	UF          string `json:"uf"`
}
