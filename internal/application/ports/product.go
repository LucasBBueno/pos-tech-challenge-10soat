package ports

import (
	"context"
	"post-tech-challenge-10soat/internal/application/core/domain"
)

type ProductRepository interface {
	ListProducts(ctx context.Context, categoryId string) ([]domain.Product, error)
	GetProductById(ctx context.Context, id string) (*domain.Product, error)
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}
