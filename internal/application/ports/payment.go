package ports

import (
	"context"
	"post-tech-challenge-10soat/internal/application/core/domain"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
}
