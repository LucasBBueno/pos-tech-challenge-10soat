package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusReceived  OrderStatus = "received"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusCompleted OrderStatus = "completed"
)

type Order struct {
	Id        uuid.UUID
	Number    int
	Status    OrderStatus
	ClientId  *uuid.UUID
	PaymentId uuid.UUID
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
