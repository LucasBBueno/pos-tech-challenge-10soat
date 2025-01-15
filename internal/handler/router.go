package handler

import (
	"log/slog"
	"os"
	"path/filepath"
	"post-tech-challenge-10soat/internal/infra/config"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	healthHandler HealthHandler,
	clientHandler ClientHandler,
	productHandler ProductHandler,
	orderHandler OrderHandler,
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

	wd, err := os.Getwd()
	if err != nil {
		panic("fail to get actual directory")
	}
	swaggerPath := filepath.Join(wd, "docs", "swagger.json")
	router.GET("/swagger.json", func(c *gin.Context) {
		c.File(swaggerPath)
	})

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.json")))

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
		product := v1.Group("/products")
		{
			product.GET("/", productHandler.ListProducts)
			product.POST("/", productHandler.CreateProduct)
			product.PUT("/:id", productHandler.UpdateProduct)
			product.DELETE("/:id", productHandler.DeleteProduct)
		}
		order := v1.Group("/orders")
		{
			order.POST("/", orderHandler.CreateOrder)
			order.GET("/", orderHandler.ListOrders)
		}
	}

	return &Router{
		router,
	}, nil
}
