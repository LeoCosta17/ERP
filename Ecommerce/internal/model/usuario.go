package model

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// UsuarioCriar é usado para criar um novo usuário, contendo os campos necessários para registro
type UsuarioCriar struct {
	Nome  string `json:"nome"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

// UsuarioLogin é usado para autenticação, contendo apenas os campos necessários para login
type UsuarioLogin struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

// UsuarioPublico inclui apenas os campos que podem ser expostos publicamente, sem informações sensíveis
type UsuarioPublico struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

// UsuarioPrivado inclui campos sensíveis como CPF e Senha, que não devem ser expostos publicamente
type UsuarioPrivado struct {
	ID           int    `json:"id"`
	Nome         string `json:"nome"`
	CPF          string `json:"cpf"`
	Telefone     string `json:"telefone"`
	Email        string `json:"email"`
	Senha        string `json:"senha"`
	CriadoEm     string `json:"criado_em"`
	AtualizadoEm string `json:"atualizado_em"`
}

// UsuarioAtualizar é usado para atualizar um usuário existente, contendo apenas os campos que podem ser atualizados
type UsuarioAtualizar struct {
	Nome     string `json:"nome,omitempty"`
	CPF      string `json:"cpf,omitempty"`
	Telefone string `json:"telefone,omitempty"`
	Email    string `json:"email,omitempty"`
}

// Validar verifica se os campos obrigatórios estão preenchidos
func (u *UsuarioCriar) Validar() error {
	var dadosInvalidos []string
	if u.Nome == "" {
		dadosInvalidos = append(dadosInvalidos, "nome")
	}
	if u.CPF == "" {
		dadosInvalidos = append(dadosInvalidos, "cpf")
	}
	if u.Email == "" {
		dadosInvalidos = append(dadosInvalidos, "email")
	}
	if u.Senha == "" {
		dadosInvalidos = append(dadosInvalidos, "senha")
	}
	if len(dadosInvalidos) > 0 {
		return errors.New("os seguintes campos são obrigatórios: " + strings.Join(dadosInvalidos, ", "))
	}
	return nil
}

// HashSenha gera um hash da senha usando bcrypt e substitui a senha original pelo hash
func (u *UsuarioCriar) HashSenha() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Senha = string(bytes)
	return nil
}

// CompararSenha compara a senha fornecida com o hash armazenado,
// retornando nil se forem iguais ou um erro se forem diferentes
func (u *UsuarioLogin) CompararSenha(senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))
}

func (u *UsuarioAtualizar) Validar() error {
	var dadosInvalidos []string
	if u.Nome == "" {
		dadosInvalidos = append(dadosInvalidos, "nome")
	}
	if u.CPF == "" {
		dadosInvalidos = append(dadosInvalidos, "cpf")
	}
	if u.Email == "" {
		dadosInvalidos = append(dadosInvalidos, "email")
	}
	if len(dadosInvalidos) > 0 {
		return errors.New("os seguintes campos são obrigatórios: " + strings.Join(dadosInvalidos, ", "))
	}
	return nil
}
