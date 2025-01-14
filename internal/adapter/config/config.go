package config

import (
	"os"
	
	"fmt"
)

type (
	Container struct {
		App  *App
		HTTP *HTTP
		DB   *DB
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}
)

func New() (*Container, error) {
	// Removido o carregamento do arquivo .env
	// Apenas verifica se as variáveis de ambiente estão definidas
	if os.Getenv("APP_ENV") == "" {
		return nil, fmt.Errorf("APP_ENV is not set")
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}
	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}
	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	return &Container{
		app,
		http,
		db,
	}, nil
}
