package product

import (
	"context"
	domain "post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type ListProducts interface {
	Execute(ctx context.Context, categoryId string) ([]domain.Product, error)
}

type ListProductsUsecase struct {
	productRepository  ports.ProductRepository
	categoryRepository ports.CategoryRepository
}

func NewListProductsUsecase(productRepository ports.ProductRepository, categoryRepository ports.CategoryRepository) ListProducts {
	return &ListProductsUsecase{
		productRepository,
		categoryRepository,
	}
}

func (s *ListProductsUsecase) Execute(ctx context.Context, categoryId string) ([]domain.Product, error) {
	var products []domain.Product
	products, err := s.productRepository.ListProducts(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	for i, product := range products {
		category, err := s.categoryRepository.GetCategoryById(ctx, product.CategoryId)
		if err != nil {
			if err == domain.ErrDataNotFound {
				return nil, err
			}
			return nil, err
		}

		products[i].Category = category
	}
	return products, nil
}
