package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

var DbConfig DatabaseConfig

func init() {
	// environment := os.Getenv("ENVIRONMENT")
	environment := "local"

	switch environment {
	case "local":
		DbConfig = DatabaseConfig{
			Username: "postgres",
			Password: "admin",
			Name:     "ps-user",
			Host:     "localhost",
			Port:     "5432",
		}
	case "prod":
		DbConfig = DatabaseConfig{
			Username: os.Getenv("DB_USUARIO_PROD"),
			Password: os.Getenv("DB_SENHA_PROD"),
			Name:     os.Getenv("DB_NOME_PROD"),
			Host:     os.Getenv("DB_HOST_PROD"),
			Port:     os.Getenv("DB_PORT_PROD"),
		}
	case "heroku":
		DbConfig = DatabaseConfig{
			Username: os.Getenv("DB_USUARIO_HEROKU"),
			Password: os.Getenv("DB_SENHA_HEROKU"),
			Name:     os.Getenv("DB_NOME_HEROKU"),
			Host:     os.Getenv("DB_HOST_HEROKU"),
			Port:     os.Getenv("DB_PORT_HEROKU"),
		}
	default:
		fmt.Println("Ambiente desconhecido. Configurações de banco de dados não encontradas.")
		os.Exit(1)
	}
}
