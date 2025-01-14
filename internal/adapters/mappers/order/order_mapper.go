package order

import (
	"github.com/google/uuid"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/utils"
	"time"
)

type OrderResponse struct {
	Id        uuid.UUID          `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Number    int                `json:"number" example:"123"`
	ClientId  uuid.UUID          `json:"client_id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Total     float64            `json:"total" example:"100.90"`
	Status    domain.OrderStatus `json:"status" example:"received"`
	CreatedAt time.Time          `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt time.Time          `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

func NewOrderResponse(order *domain.Order) *OrderResponse {
	orderResponse := &OrderResponse{
		Id:        utils.StringToUuid(order.Id),
		Number:    order.Number,
		Total:     order.Total,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
	if order.ClientId != "" {
		orderResponse.ClientId = utils.StringToUuid(order.ClientId)
	}
	return orderResponse
}

type ListOrdersResponse struct {
	ReceivedOrders  []OrderResponse `json:"received_orders"`
	PreparingOrders []OrderResponse `json:"preparing_orders"`
	ReadyOrders     []OrderResponse `json:"ready_orders"`
	CompletedOrders []OrderResponse `json:"completed_orders"`
}

func NewListOrdersResponse(listOrders *domain.ListOrders) *ListOrdersResponse {
	var receivedOrdersResponse []OrderResponse
	var preparingOrdersResponse []OrderResponse
	var readyOrdersResponse []OrderResponse
	var completedOrdersResponse []OrderResponse
	for _, order := range listOrders.ReceivedOrders {
		receivedOrdersResponse = append(receivedOrdersResponse, *NewOrderResponse(&order))
	}
	for _, order := range listOrders.PreparingOrders {
		preparingOrdersResponse = append(preparingOrdersResponse, *NewOrderResponse(&order))
	}
	for _, order := range listOrders.ReadyOrders {
		readyOrdersResponse = append(readyOrdersResponse, *NewOrderResponse(&order))
	}
	for _, order := range listOrders.CompletedOrders {
		completedOrdersResponse = append(completedOrdersResponse, *NewOrderResponse(&order))
	}
	return &ListOrdersResponse{
		ReceivedOrders:  receivedOrdersResponse,
		PreparingOrders: preparingOrdersResponse,
		ReadyOrders:     readyOrdersResponse,
		CompletedOrders: completedOrdersResponse,
	}
}
