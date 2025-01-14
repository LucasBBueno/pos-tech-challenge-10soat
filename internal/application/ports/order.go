package ports

import (
	"context"
	"post-tech-challenge-10soat/internal/application/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context, limit uint64) (*domain.ListOrders, error)
}
