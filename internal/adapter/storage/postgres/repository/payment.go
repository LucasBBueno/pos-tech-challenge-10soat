package repository

import (
	"context"
	"post-tech-challenge-10soat/internal/adapter/storage/postgres"
	"post-tech-challenge-10soat/internal/core/domain"
)

type PaymentRepository struct {
	db *postgres.DB
}

func NewPaymentRepository(db *postgres.DB) *PaymentRepository {
	return &PaymentRepository{
		db,
	}
}

func (repository *PaymentRepository) CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {
	query := repository.db.QueryBuilder.Insert("payments").
		Columns("type", "provider").
		Values(payment.Type, payment.Provider).
		Suffix("RETURNING *")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&payment.Id,
		&payment.Type,
		&payment.Provider,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
