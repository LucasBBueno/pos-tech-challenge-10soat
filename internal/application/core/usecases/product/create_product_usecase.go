package product

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type CreateProduct interface {
	Execute(ctx context.Context, product *domain.Product) (*domain.Product, error)
}

type CreateProductUsecase struct {
	productRepository  ports.ProductRepository
	categoryRepository ports.CategoryRepository
}

func NewCreateProductUsecase(productRepository ports.ProductRepository, categoryRepository ports.CategoryRepository) CreateProduct {
	return &CreateProductUsecase{
		productRepository,
		categoryRepository,
	}
}

func (s *CreateProductUsecase) Execute(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	category, err := s.categoryRepository.GetCategoryById(ctx, product.CategoryId)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create product for this category - %s", err.Error())
	}
	product.Category = category
	product, err = s.productRepository.CreateProduct(ctx, product)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return product, nil
}
