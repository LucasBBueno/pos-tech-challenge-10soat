package port

import (
	"context"
	"post-tech-challenge-10soat/internal/core/domain"
)

type OrderProductRepository interface {
	CreateOrderProduct(ctx context.Context, payment *domain.OrderProduct) (*domain.OrderProduct, error)
}
