package port

import (
	"context"
	"post-tech-challenge-10soat/internal/core/domain"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
}

type PaymentService interface {
	Checkout(ctx context.Context, createPayment *domain.CreatePayment) (*domain.Payment, error)
}
