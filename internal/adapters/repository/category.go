package repository

import (
	"context"
	domain "post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/infra/storage/postgres"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

type CategoryRepository struct {
	db *postgres.DB
}

func NewCategoryRepository(db *postgres.DB) *CategoryRepository {
	return &CategoryRepository{
		db,
	}
}

func (cr *CategoryRepository) GetCategoryById(ctx context.Context, id string) (*domain.Category, error) {
	var category domain.Category
	query := cr.db.QueryBuilder.Select("*").
		From("categories").
		Where(sq.Eq{"id": id}).
		Limit(1)
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
