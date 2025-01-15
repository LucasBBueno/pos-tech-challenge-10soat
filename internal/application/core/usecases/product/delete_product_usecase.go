package product

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type DeleteProduct interface {
	Execute(ctx context.Context, id string) error
}

type DeleteProductUsecase struct {
	productRepository ports.ProductRepository
}

func NewDeleteProductUsecase(productRepository ports.ProductRepository) DeleteProduct {
	return &DeleteProductUsecase{
		productRepository,
	}
}

func (s *DeleteProductUsecase) Execute(ctx context.Context, id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid product id")
	}
	_, err = s.productRepository.GetProductById(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return fmt.Errorf("cannot delete product for this identifier - %s", err.Error())
	}
	return s.productRepository.DeleteProduct(ctx, id)
}
