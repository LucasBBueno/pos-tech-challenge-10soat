package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (handler *HealthHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Application is healthy",
	})
}
