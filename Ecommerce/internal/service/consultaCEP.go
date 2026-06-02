package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/leona/ecommerce/internal/model"
)

func ConsultaCEP(cep string) (*model.EnderecoConsultaCEP, error) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Erro ao consultar cep!")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var endereco model.EnderecoConsultaCEP
	if err := json.Unmarshal(body, &endereco); err != nil {
		return nil, err
	}

	return &endereco, nil
}
