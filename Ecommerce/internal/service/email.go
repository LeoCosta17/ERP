package service

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/leona/ecommerce/configs"
	"github.com/resend/resend-go/v3"
)

func EnviarEmail(nome, token, email string) error {

	fmt.Printf("Enviando email para: %s\n", email)

	apikey := configs.GetString("EMAIL_API_KEY", "re_NGPkJgxY_4Kt6c72Dg6ga7cXTfoGzV7s4")

	client := resend.NewClient(apikey)

	template, err := template.ParseFiles("C:\\Users\\leona\\OneDrive\\Área de Trabalho\\Ecommerce\\internal\\template\\modeloEmailValidacaoConta.html")
	if err != nil {
		return err
	}

	var emailContent bytes.Buffer
	data := map[string]interface{}{
		"Nome":  nome,
		"Token": token,
	}
	err = template.Execute(&emailContent, data)
	if err != nil {
		return err
	}

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{email},
		Subject: "Ative sua Conta",
		Html:    emailContent.String(),
	}

	response, err := client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Printf("email sent: %s", response.Id)

	return nil
}
