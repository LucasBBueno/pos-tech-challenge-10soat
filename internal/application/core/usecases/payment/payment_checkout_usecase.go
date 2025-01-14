package payment

import (
	"context"
	"fmt"
	domain2 "post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type PaymentCheckout interface {
	Execute(ctx context.Context, createPayment *domain2.CreatePayment) (*domain2.Payment, error)
}

type PaymentCheckoutUsecase struct {
	paymentRepository ports.PaymentRepository
}

func NewPaymentCheckoutUsecase(paymentRepository ports.PaymentRepository) *PaymentCheckoutUsecase {
	return &PaymentCheckoutUsecase{
		paymentRepository,
	}
}

func (s *PaymentCheckoutUsecase) Execute(ctx context.Context, createPayment *domain2.CreatePayment) (*domain2.Payment, error) {
	//TODO: Checkout no fornecedor MP
	paymentInfo := domain2.Payment{
		Type:     createPayment.Type,
		Provider: createPayment.Provider,
	}
	payment, err := s.paymentRepository.CreatePayment(ctx, &paymentInfo)
	if err != nil {
		if err == domain2.ErrConflictingData {
			return nil, err
		}
		return nil, fmt.Errorf("failed to make payment - %s", err.Error())
	}
	return payment, nil
}
