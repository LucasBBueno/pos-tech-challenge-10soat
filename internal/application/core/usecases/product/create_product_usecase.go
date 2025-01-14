package product

import (
	"context"
	"fmt"
	domain2 "post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type CreateProductUsecase struct {
	productRepository  ports.ProductRepository
	categoryRepository ports.CategoryRepository
}

func NewCreateProductUsecase(productRepository ports.ProductRepository, categoryRepository ports.CategoryRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		productRepository,
		categoryRepository,
	}
}

func (s *CreateProductUsecase) Execute(ctx context.Context, product *domain2.Product) (*domain2.Product, error) {
	category, err := s.categoryRepository.GetCategoryById(ctx, product.CategoryId)
	if err != nil {
		if err == domain2.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create product for this category - %s", err.Error())
	}
	product.Category = category
	product, err = s.productRepository.CreateProduct(ctx, product)
	if err != nil {
		if err == domain2.ErrConflictingData {
			return nil, err
		}
		return nil, domain2.ErrInternal
	}
	return product, nil
}
