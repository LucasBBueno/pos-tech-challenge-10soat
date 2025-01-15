package ports

import (
	"context"
	"post-tech-challenge-10soat/internal/application/core/domain"
)

type CategoryRepository interface {
	GetCategoryById(ctx context.Context, categoryId string) (*domain.Category, error)
}
