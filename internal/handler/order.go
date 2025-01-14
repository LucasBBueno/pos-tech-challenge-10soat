package handler

import (
	"github.com/gin-gonic/gin"
	order2 "post-tech-challenge-10soat/internal/adapters/mappers/order"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/core/usecases/order"
)

type OrderHandler struct {
	createOrder order.CreateOrder
	listOrder   order.ListOrders
}

func NewOrderHandler(service order.CreateOrder, listOrders order.ListOrders) *OrderHandler {
	return &OrderHandler{
		service,
		listOrders,
	}
}

type orderProductRequest struct {
	ProductID   string `json:"product_id" binding:"required,min=1" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Quantity    int    `json:"quantity" binding:"required,number" example:"1"`
	Observation string `json:"observation" binding:"omitempty" example:"Lanche com batata"`
}

type createOrderRequest struct {
	ClientId string                `json:"client_id" binding:"omitempty" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Products []orderProductRequest `json:"products" binding:"required"`
}

// CreateOrder godoc
//
//	@Summary		Criar um novo pedido
//	@Description	Cria um novo pedido processando o pagamento
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			createOrderRequest	body		createOrderRequest	true	"Criar ordem body"
//	@Success		200					{object}	orderResponse		"Ordem criada"
//	@Failure		400					{object}	errorResponse		"Erro de validação"
//	@Failure		500					{object}	errorResponse		"Erro interno"
//	@Router			/orders [post]
//	@Security		BearerAuth
func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var request createOrderRequest
	var products []domain.CreateOrderProduct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}
	for _, product := range request.Products {
		products = append(products, domain.CreateOrderProduct{
			ProductId:   product.ProductID,
			Quantity:    product.Quantity,
			Observation: product.Observation,
		})
	}
	oderInfo := domain.CreateOrder{
		ClientId: request.ClientId,
		Products: products,
	}
	o, err := h.createOrder.Execute(ctx, &oderInfo)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := order2.NewOrderResponse(o)
	handleSuccess(ctx, response)
}

type listOrdersRequest struct {
	Limit uint64 `form:"limit" binding:"required,min=1" example:5""`
}

// ListOrders godoc
//
//	@Summary		Lista os pedidos
//	@Description	Lista os pedidos separados por status com limite de consulta
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int			false	"Limite de pedidos"
//	@Success		200			{object}	listOrdersResponse			"Pedidos listados"
//	@Failure		400			{object}	errorResponse	"Erro de validação"
//	@Failure		500			{object}	errorResponse	"Erro interno"
//	@Router			/orders [get]
func (h *OrderHandler) ListOrders(ctx *gin.Context) {
	var request listOrdersRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		validationError(ctx, err)
		return
	}
	listOrders, err := h.listOrder.Execute(ctx, request.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := order2.NewListOrdersResponse(listOrders)
	handleSuccess(ctx, response)
}
