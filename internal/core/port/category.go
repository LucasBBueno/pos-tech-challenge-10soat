package port

import (
	"context"
	"post-tech-challenge-10soat/internal/core/domain"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	GetCategoryById(ctx context.Context, categoryId uuid.UUID) (*domain.Category, error)
}
