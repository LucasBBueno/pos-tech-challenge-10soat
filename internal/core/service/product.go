package service

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/core/domain"
	"post-tech-challenge-10soat/internal/core/port"

	"github.com/google/uuid"
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
		category, err := service.categoryRepository.GetCategoryById(ctx, product.CategoryId.String())
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

func (service *ProductService) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	category, err := service.categoryRepository.GetCategoryById(ctx, product.CategoryId.String())
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create product for this category - %s", err.Error())
	}
	product.Category = category
	product, err = service.productRepository.CreateProduct(ctx, product)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return product, nil
}

func (service *ProductService) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	existingProduct, err := service.productRepository.GetProductById(ctx, product.Id.String())
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot find product to update - %s", err.Error())
	}
	emptyData := uuid.Validate(product.CategoryId.String()) != nil &&
		product.Name == "" &&
		product.Value == 0
	sameData := existingProduct.CategoryId.String() == product.CategoryId.String() &&
		existingProduct.Name == product.Name &&
		existingProduct.Value == product.Value
	if emptyData || sameData {
		return nil, domain.ErrNoUpdatedData
	}
	if uuid.Validate(product.CategoryId.String()) != nil {
		product.CategoryId = existingProduct.CategoryId
	}
	category, err := service.categoryRepository.GetCategoryById(ctx, product.CategoryId.String())
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot update product for this category - %s", err.Error())
	}
	product.Category = category
	_, err = service.productRepository.UpdateProduct(ctx, product)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return product, nil
}

func (service *ProductService) DeleteProduct(ctx context.Context, id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid product id")
	}
	_, err = service.productRepository.GetProductById(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return fmt.Errorf("cannot delete product for this identifier - %s", err.Error())
	}
	return service.productRepository.DeleteProduct(ctx, id)
}
