package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewController struct {
}

func (c *ViewController) RenderizarLoginPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"web/template/pages/login.html",
		"web/template/components/loginForm.html",
	)

	if err != nil {
		fmt.Printf("Erro ao renderizar página de login: %v\n", err)
		http.Error(w, "Erro ao renderizar página de login", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Printf("Erro ao renderizar página de login: %v\n", err)
		http.Error(w, "Erro ao renderizar página de login", http.StatusInternalServerError)
		return
	}
}

func (c *ViewController) RenderizarDashboardPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/template/pages/dashboard.html")
	if err != nil {
		fmt.Printf("Erro ao renderizar dashboard: %v\n", err)
		http.Error(w, "Erro ao renderizar dashboard", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Printf("Erro ao renderizar dashboard: %v\n", err)
		http.Error(w, "Erro ao renderizar dashboard", http.StatusInternalServerError)
		return
	}
}
