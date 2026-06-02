package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/leona/ecommerce/configs"
	"github.com/leona/ecommerce/db"
	"github.com/leona/ecommerce/internal/controller"
	"github.com/leona/ecommerce/internal/repository"
	"github.com/leona/ecommerce/internal/router"
	"github.com/leona/ecommerce/internal/service"
	// Chi router
)

func main() {
	app := &application{
		API_port:           fmt.Sprintf(":%d", configs.GetInt("API_PORT", 8080)),
		API_maxHeaderBytes: configs.GetInt("API_MAX_HEADER_BYTES", 1<<20), // 1 MB
		API_readTimeout:    configs.GetDuration("API_READ_TIMEOUT", 10*time.Second),
		API_writeTimeout:   configs.GetDuration("API_WRITE_TIMEOUT", 10*time.Second),
	}

	database := &db.BancoDados{
		ConnectionString: configs.GetString("DB_CONNECTION_STRING", "admin:12345@tcp(localhost:3307)/ecommerce"),
		Driver:           configs.GetString("DB_DRIVER", "mysql"),
		MaxOpenConns:     configs.GetInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns:     configs.GetInt("DB_MAX_IDLE_CONNS", 25),
		MaxIdleTime:      configs.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
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
