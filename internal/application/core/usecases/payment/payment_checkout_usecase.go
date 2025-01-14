package payment

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type PaymentCheckout interface {
	Execute(ctx context.Context, createPayment *domain.CreatePayment) (*domain.Payment, error)
}

type PaymentCheckoutUsecase struct {
	pr ports.PaymentRepository
}

func NewPaymentCheckoutUsecase(paymentRepository ports.PaymentRepository) PaymentCheckout {
	return &PaymentCheckoutUsecase{
		paymentRepository,
	}
}

func (s *PaymentCheckoutUsecase) Execute(ctx context.Context, createPayment *domain.CreatePayment) (*domain.Payment, error) {
	//TODO: Checkout no fornecedor MP
	paymentInfo := domain.Payment{
		Type:     createPayment.Type,
		Provider: createPayment.Provider,
	}
	payment, err := s.pr.CreatePayment(ctx, &paymentInfo)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, fmt.Errorf("failed to make payment - %s", err.Error())
	}
	return payment, nil
}
