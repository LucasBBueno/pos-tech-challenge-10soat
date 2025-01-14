package product

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	domain2 "post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type UpdateProductUsecase struct {
	productRepository  ports.ProductRepository
	categoryRepository ports.CategoryRepository
}

func NewUpdateProductUsecase(productRepository ports.ProductRepository, categoryRepository ports.CategoryRepository) *UpdateProductUsecase {
	return &UpdateProductUsecase{
		productRepository,
		categoryRepository,
	}
}

func (s *UpdateProductUsecase) Execute(ctx context.Context, product *domain2.Product) (*domain2.Product, error) {
	existingProduct, err := s.productRepository.GetProductById(ctx, product.Id)
	if err != nil {
		if err == domain2.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot find product to update - %s", err.Error())
	}
	emptyData := uuid.Validate(product.CategoryId) != nil &&
		product.Name == "" &&
		product.Value == 0
	sameData := existingProduct.CategoryId == product.CategoryId &&
		existingProduct.Name == product.Name &&
		existingProduct.Value == product.Value
	if emptyData || sameData {
		return nil, domain2.ErrNoUpdatedData
	}
	if uuid.Validate(product.CategoryId) != nil {
		product.CategoryId = existingProduct.CategoryId
	}
	category, err := s.categoryRepository.GetCategoryById(ctx, product.CategoryId)
	if err != nil {
		if err == domain2.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot update product for this category - %s", err.Error())
	}
	product.Category = category
	_, err = s.productRepository.UpdateProduct(ctx, product)
	if err != nil {
		if err == domain2.ErrConflictingData {
			return nil, err
		}
		return nil, domain2.ErrInternal
	}
	return product, nil
}
