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

// CreateClient godoc
//
//	@Summary     Registra um novo cliente
//	@Description Registra um novo cliente com nome e e-mail
//	@Tags        Clients
//	@Accept      json
//	@Produce		json
//	@Param	    createClientRequest	body createClientRequest true "Registrar novo cliente request"
//	@Success		200	{object} clientResponse	"Cliente registrado"
//	@Failure		400	{object} errorResponse	"Erro de validação"
//	@Router		/clients [post]
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

// GetClientByCpf godoc
//
//	    @Summary     Busca um cliente
//	    @Description buscar um cliente pelo Cpf
//	    @Tags        Clients
//	    @Accept      json
//	    @Produce		json
//		   @Param	    cpf	path		string				true	"CPF"
//	    @Success		200	{object}    clientResponse	"Cliente"
//	    @Failure		400	{object}    errorResponse	"Erro de validação"
//		   @Failure		404	{object}	errorResponse   "Cliente nao encontrado"
//	    @Router		/clients/{cpf} [get]
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
