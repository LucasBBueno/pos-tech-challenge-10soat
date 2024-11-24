package main

import (
	"fmt"
	"log/slog"
	"os"

	"post-tech-challenge-10soat/internal/adapter/config"
	"post-tech-challenge-10soat/internal/adapter/logger"
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

	healthHandler := handler.NewHealthHandler()
	router, err := handler.NewRouter(
		config.HTTP,
		*healthHandler,
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
