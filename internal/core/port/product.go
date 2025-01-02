package port

import (
	"context"
	"post-tech-challenge-10soat/internal/core/domain"
)

type ProductRepository interface {
	ListProducts(ctx context.Context, categoryId string) ([]domain.Product, error)
}

type ProductService interface {
	ListProducts(ctx context.Context, categoryId string) ([]domain.Product, error)
}
