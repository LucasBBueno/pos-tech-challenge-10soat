package service

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/core/domain"
	"post-tech-challenge-10soat/internal/core/port"
)

type PaymentService struct {
	paymentRepository port.PaymentRepository
}

func NewPaymentService(paymentRepository port.PaymentRepository) *PaymentService {
	return &PaymentService{
		paymentRepository,
	}
}

func (service *PaymentService) Checkout(ctx context.Context, createPayment *domain.CreatePayment) (*domain.Payment, error) {
	//TODO: Checkout no fornecedor MP
	paymentInfo := domain.Payment{
		Type:     createPayment.Type,
		Provider: createPayment.Provider,
	}
	payment, err := service.paymentRepository.CreatePayment(ctx, &paymentInfo)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, fmt.Errorf("failed to make payment - %s", err.Error())
	}
	return payment, nil
}
