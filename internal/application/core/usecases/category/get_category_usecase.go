package category

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type GetCategory interface {
	Execute(ctx context.Context, id string) (*domain.Category, error)
}

type GetCategoryUsecase struct {
	repository ports.CategoryRepository
}

func NewGetCategoryUsecase(repository ports.CategoryRepository) GetCategory {
	return &GetCategoryUsecase{
		repository,
	}
}

func (s *GetCategoryUsecase) Execute(ctx context.Context, id string) (*domain.Category, error) {
	c, err := s.repository.GetCategoryById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get category by id - %s", err.Error())
	}
	return c, nil
}
