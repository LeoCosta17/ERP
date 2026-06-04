package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UsuarioCriar struct {
	IDEmpresa uint64 `json:"id_empresa"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
}

type UsuarioLogin struct {
	IDEmpresa uint64 `json:"id_empresa"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
}

type UsuarioBasico struct {
	ID        uint64 `json:"id"`
	IDEmpresa uint64 `json:"id_empresa"`
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
}

type UsuarioCompleto struct {
	ID              uint64 `json:"id"`
	IDEmpresa       uint64 `json:"id_empresa"`
	Nome            string `json:"nome"`
	Cpf             string `json:"cpf"`
	Telefone        string `json:"telefone"`
	Email           string `json:"email"`
	Ativo           bool   `json:"ativo"`
	DataCriacao     string `json:"data_criacao"`
	DataAtualizacao string `json:"data_atualizacao"`
}

func (u UsuarioCriar) Validar() error {
	var erros []error
	if u.IDEmpresa == 0 {
		erros = append(erros, errors.New("o id da empresa não foi informado"))
	}
	if u.Nome == "" {
		erros = append(erros, errors.New("o nome não foi informado"))
	}
	if u.Email == "" {
		erros = append(erros, errors.New("o email não foi informado"))
	}
	if u.Senha == "" {
		erros = append(erros, errors.New("a senha não foi informada"))
	}
	if len(erros) > 0 {
		return errors.Join(erros...)
	}
	return nil
}

func (u UsuarioCriar) HashSenha() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Senha = string(hash)
	return nil
}

func (u UsuarioLogin) Validar() error {
	var erros []error
	if u.Email == "" {
		erros = append(erros, errors.New("o email não foi informado"))
	}
	if u.Senha == "" {
		erros = append(erros, errors.New("a senha não foi informada"))
	}
	if len(erros) > 0 {
		return errors.Join(erros...)
	}
	return nil
}

func (u UsuarioLogin) ValidarSenha(senhaDB string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(senhaDB), []byte(u.Senha)); err != nil {
		return errors.New("senha inválida")
	}
	return nil
}
