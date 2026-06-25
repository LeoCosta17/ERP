package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
	senha := "123456"

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	if err != nil {
		fmt.Println("ERRO:", err)
	} else {
		fmt.Println("SUCESSO")
	}
}
