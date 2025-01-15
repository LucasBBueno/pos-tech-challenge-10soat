package repository

import (
	"context"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/infra/storage/postgres"
)

type OrderProductRepository struct {
	db *postgres.DB
}

func NewOrderProductRepository(db *postgres.DB) *OrderProductRepository {
	return &OrderProductRepository{
		db,
	}
}

func (repository *OrderProductRepository) CreateOrderProduct(ctx context.Context, orderProduct *domain.OrderProduct) (*domain.OrderProduct, error) {
	query := repository.db.QueryBuilder.Insert("order_products").
		Columns("order_id", "product_id", "quantity", "sub_total", "observation").
		Values(orderProduct.OrderId, orderProduct.ProductId, orderProduct.Quantity, orderProduct.SubTotal, orderProduct.Observation).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&orderProduct.Id,
		&orderProduct.OrderId,
		&orderProduct.ProductId,
		&orderProduct.Quantity,
		&orderProduct.SubTotal,
		&orderProduct.Observation,
		&orderProduct.CreatedAt,
		&orderProduct.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return orderProduct, nil
}
