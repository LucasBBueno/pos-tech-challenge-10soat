package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"post-tech-challenge-10soat/internal/adapter/storage/postgres"
	"post-tech-challenge-10soat/internal/core/domain"
)

type ProductRepository struct {
	db *postgres.DB
}

func NewProductRepository(db *postgres.DB) *ProductRepository {
	return &ProductRepository{
		db,
	}
}

func (repository *ProductRepository) ListProducts(ctx context.Context, categoryId string) ([]domain.Product, error) {
	var product domain.Product
	var products []domain.Product
	query := repository.db.QueryBuilder.Select("*").
		From("products").
		OrderBy("created_at")

	if categoryId != "" {
		err := uuid.Validate(categoryId)
		if err != nil {
			return []domain.Product{}, fmt.Errorf("invalid category")
		}
		query = query.Where(sq.Eq{"category_id": categoryId})
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return []domain.Product{}, err
	}
	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return []domain.Product{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Image,
			&product.Value,
			&product.CategoryId,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return []domain.Product{}, err
		}

		products = append(products, product)
	}
	return products, nil
}
