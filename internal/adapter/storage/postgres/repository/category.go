package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"

	"post-tech-challenge-10soat/internal/adapter/storage/postgres"
	"post-tech-challenge-10soat/internal/core/domain"
)

type CategoryRepository struct {
	db *postgres.DB
}

func NewCategoryRepository(db *postgres.DB) *CategoryRepository {
	return &CategoryRepository{
		db,
	}
}

func (cr *CategoryRepository) GetCategoryById(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	var category domain.Category
	query := cr.db.QueryBuilder.Select("*").
		From("categories").
		Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}
	return &category, nil
}
