package handler

import (
	"post-tech-challenge-10soat/internal/core/domain"
	"post-tech-challenge-10soat/internal/core/port"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	service port.ClientService
}

func NewClientHandler(service port.ClientService) *ClientHandler {
	return &ClientHandler{
		service,
	}
}

type createClientRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required" example:"john-doe@email.com"`
}

func (handler *ClientHandler) CreateClient(ctx *gin.Context) {
	var request createClientRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}
	client := domain.Client{
		Name:  request.Name,
		Email: request.Email,
	}
	_, err := handler.service.CreateClient(ctx, &client)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := newClientReponse(&client)
	handleSuccess(ctx, response)
}

type getClientByCpfRequest struct {
	Cpf string `uri:"cpf" binding:"required,min=1" example:"12345678010"`
}

func (handler *ClientHandler) GetClientByCpf(ctx *gin.Context) {
	var request getClientByCpfRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		validationError(ctx, err)
		return
	}
	client, err := handler.service.GetClientByCpf(ctx, request.Cpf)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := newClientReponse(client)
	handleSuccess(ctx, response)
}
