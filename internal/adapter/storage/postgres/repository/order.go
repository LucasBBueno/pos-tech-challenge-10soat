package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"post-tech-challenge-10soat/internal/adapter/storage/postgres"
	"post-tech-challenge-10soat/internal/core/domain"
)

type OrderRepository struct {
	db *postgres.DB
}

func NewOrderRepository(db *postgres.DB) *OrderRepository {
	return &OrderRepository{
		db,
	}
}

func (repository *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	query := repository.db.QueryBuilder.Insert("orders").
		Columns("status", "client_id", "payment_id", "total").
		Values(order.Status, order.ClientId, order.PaymentId, order.Total).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&order.Id,
		&order.Number,
		&order.Status,
		&order.ClientId,
		&order.PaymentId,
		&order.Total,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repository *OrderRepository) DeleteOrder(ctx context.Context, id string) error {
	query := repository.db.QueryBuilder.Delete("orders").
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
