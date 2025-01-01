package handler

import (
	"log/slog"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"

	"post-tech-challenge-10soat/internal/adapter/config"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	healthHandler HealthHandler,
	clientHandler ClientHandler,
) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("/v1")
	{
		health := v1.Group("/health")
		{
			health.GET("/", healthHandler.HealthCheck)
		}
		client := v1.Group("/clients")
		{
			client.POST("/", clientHandler.CreateClient)
			client.GET("/:cpf", clientHandler.GetClientByCpf)
		}
	}
	return &Router{
		router,
	}, nil
}
