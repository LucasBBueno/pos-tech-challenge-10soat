package repository

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/infra/storage/postgres"
	"post-tech-challenge-10soat/internal/utils"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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

func (repository *ProductRepository) GetProductById(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product
	query := repository.db.QueryBuilder.Select("*").
		From("products").
		Where(sq.Eq{"id": id}).
		Limit(1)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
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
		return nil, err
	}
	return &product, nil
}

func (repository *ProductRepository) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	query := repository.db.QueryBuilder.Insert("products").
		Columns("name", "description", "image", "value", "category_id").
		Values(product.Name, product.Description, product.Image, product.Value, product.CategoryId).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
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
		return nil, err
	}
	return product, nil
}

func (repository *ProductRepository) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	name := utils.NullString(product.Name)
	description := utils.NullString(product.Description)
	image := utils.NullString(product.Image)
	query := repository.db.QueryBuilder.Update("products").
		Set("name", sq.Expr("COALESCE(?, name)", name)).
		Set("description", sq.Expr("COALESCE(?, description)", description)).
		Set("image", sq.Expr("COALESCE(?, image)", image)).
		Set("value", sq.Expr("COALESCE(?, value)", product.Value)).
		Set("category_id", sq.Expr("COALESCE(?, category_id)", product.CategoryId)).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": product.Id}).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
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
		return nil, err
	}
	return product, nil
}

func (repository *ProductRepository) DeleteProduct(ctx context.Context, id string) error {
	query := repository.db.QueryBuilder.Delete("products").
		Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = repository.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
