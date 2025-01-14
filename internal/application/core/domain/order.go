package domain

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusReceived  OrderStatus = "received"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusCompleted OrderStatus = "completed"
)

type Order struct {
	Id        string
	Number    int
	Status    OrderStatus
	ClientId  string
	PaymentId string
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateOrderProduct struct {
	ProductId   string
	Quantity    int
	Observation string
	SubTotal    float64
}

type CreateOrder struct {
	Status   OrderStatus
	ClientId string
	Total    float64
	Products []CreateOrderProduct
}

type ListOrders struct {
	ReceivedOrders  []Order
	PreparingOrders []Order
	ReadyOrders     []Order
	CompletedOrders []Order
}
