package handler

import (
	"post-tech-challenge-10soat/internal/core/port"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service port.ProductService
}

func NewProductHandler(service port.ProductService) *ProductHandler {
	return &ProductHandler{
		service,
	}
}

type listProductsRequest struct {
	CategoryID string `form:"category_id" binding:"omitempty,min=1" example:ed6ac028-8016-4cbd-aeee-c3a155cdb2a4""`
}

// ListProducts godoc
//
//	@Summary		Lista os produtos
//	@Description	Lista os produtos podendo buscar por categoria
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			category_id	query		string			false	"Id da categoria"
//	@Success		200			{array}	productResponse			"Produtos listados"
//	@Failure		400			{object}	errorResponse	"Erro de validação"
//	@Failure		500			{object}	errorResponse	"Erro interno"
//	@Router			/products [get]
func (handler *ProductHandler) ListProducts(ctx *gin.Context) {
	var request listProductsRequest
	var productsList []productResponse
	if err := ctx.ShouldBindQuery(&request); err != nil {
		validationError(ctx, err)
		return
	}
	products, err := handler.service.ListProducts(ctx, request.CategoryID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, product := range products {
		productsList = append(productsList, newProductResponse(&product))
	}
	handleSuccess(ctx, productsList)
}
