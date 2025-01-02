package service

import (
	"context"
	"post-tech-challenge-10soat/internal/core/domain"
	"post-tech-challenge-10soat/internal/core/port"
)

type ProductService struct {
	productRepository  port.ProductRepository
	categoryRepository port.CategoryRepository
}

func NewProductService(productRepository port.ProductRepository, categoryRepository port.CategoryRepository) *ProductService {
	return &ProductService{
		productRepository,
		categoryRepository,
	}
}

func (service *ProductService) ListProducts(ctx context.Context, categoryId string) ([]domain.Product, error) {
	var products []domain.Product
	products, err := service.productRepository.ListProducts(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	for i, product := range products {
		category, err := service.categoryRepository.GetCategoryById(ctx, product.CategoryId)
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
