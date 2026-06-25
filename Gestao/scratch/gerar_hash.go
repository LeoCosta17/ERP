package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "123456"
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("ERRO:", err)
		return
	}
	fmt.Println("HASH GERADO:", string(hash))
}
