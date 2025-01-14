package ports

import (
	"context"
	"post-tech-challenge-10soat/internal/application/core/domain"
)

type OrderProductRepository interface {
	CreateOrderProduct(ctx context.Context, payment *domain.OrderProduct) (*domain.OrderProduct, error)
}
