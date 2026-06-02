package main

import (
	"fmt"
	"gestao/config"
	"gestao/db"
	"gestao/internal/controller"
	"gestao/internal/repository"
	"gestao/internal/router"
	"gestao/internal/service"
	"time"
)

func main() {
	app := &application{
		API_port:           fmt.Sprintf(":%d", config.GetInt("API_PORT", 8080)),
		API_maxHeaderBytes: config.GetInt("API_MAX_HEADER_BYTES", 1<<20), // 1 MB
		API_readTimeout:    config.GetDuration("API_READ_TIMEOUT", 10*time.Second),
		API_writeTimeout:   config.GetDuration("API_WRITE_TIMEOUT", 10*time.Second),
	}

	database := &db.BancoDados{
		ConnectionString: config.GetString("DB_CONNECTION_STRING", "admin:12345@tcp(localhost:3307)/ecommerce"),
		Driver:           config.GetString("DB_DRIVER", "mysql"),
		MaxOpenConns:     config.GetInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns:     config.GetInt("DB_MAX_IDLE_CONNS", 25),
		MaxIdleTime:      config.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
	}

	db, err := database.Conectar()
	if err != nil {
		fmt.Printf("Erro ao conectar ao banco de dados: %v\n", err)
		return
	}
	defer db.Close()

	repositorio := repository.NewRepository(db)
	service := service.NewService(repositorio, db)
	controller := controller.NewController(service)

	r := router.CarregarRotas(controller)

	if err := app.iniciarApp(r); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
