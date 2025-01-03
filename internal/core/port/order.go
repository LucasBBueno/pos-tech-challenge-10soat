package port

import (
	"context"
	"post-tech-challenge-10soat/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type OrderService interface {
	CreateOrder(ctx context.Context, createOrder *domain.CreateOrder) (*domain.Order, error)
}
