package repository

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/infra/storage/postgres"

	sq "github.com/Masterminds/squirrel"
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

func (repository *OrderRepository) ListOrders(ctx context.Context, limit uint64) (*domain.ListOrders, error) {
	receivedOrders, err := repository.getOrdersByStatus(ctx, domain.OrderStatusReceived, limit)
	if err != nil {
		return nil, err
	}
	preparingOrders, err := repository.getOrdersByStatus(ctx, domain.OrderStatusPreparing, limit)
	if err != nil {
		return nil, err
	}
	readyOrders, err := repository.getOrdersByStatus(ctx, domain.OrderStatusReady, limit)
	if err != nil {
		return nil, err
	}
	completedOrders, err := repository.getOrdersByStatus(ctx, domain.OrderStatusCompleted, limit)
	if err != nil {
		return nil, err
	}
	return &domain.ListOrders{
		ReceivedOrders:  receivedOrders,
		PreparingOrders: preparingOrders,
		ReadyOrders:     readyOrders,
		CompletedOrders: completedOrders,
	}, nil
}

func (repository *OrderRepository) getOrdersByStatus(ctx context.Context, status domain.OrderStatus, limit uint64) ([]domain.Order, error) {
	var order domain.Order
	var orders []domain.Order
	query := repository.db.QueryBuilder.Select("*").
		From("orders").
		Where(sq.Eq{"status": status}).
		OrderBy("created_at").
		Limit(limit)
	sql, args, err := query.ToSql()
	if err != nil {
		return []domain.Order{}, fmt.Errorf("failed to get orders - %s", err.Error())
	}
	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return []domain.Order{}, fmt.Errorf("failed to get orders - %s", err.Error())
	}
	for rows.Next() {
		err := rows.Scan(
			&order.Id,
			&order.Number,
			&order.Status,
			&order.ClientId,
			&order.PaymentId,
			&order.Total,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err == nil {
			orders = append(orders, order)
		}
	}
	return orders, nil
}
