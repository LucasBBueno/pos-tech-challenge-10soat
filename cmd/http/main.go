package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"post-tech-challenge-10soat/internal/adapter/config"
	"post-tech-challenge-10soat/internal/adapter/logger"
	"post-tech-challenge-10soat/internal/adapter/storage/postgres"
	"post-tech-challenge-10soat/internal/adapter/storage/postgres/repository"
	"post-tech-challenge-10soat/internal/core/service"
	"post-tech-challenge-10soat/internal/handler"
)

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}
	logger.Set(config.App)
	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	healthHandler := handler.NewHealthHandler()

	clientRepository := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepository)
	clientHandler := handler.NewClientHandler(clientService)

	router, err := handler.NewRouter(
		config.HTTP,
		*healthHandler,
		*clientHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	listenAddress := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddress)
	err = router.Run(listenAddress)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
