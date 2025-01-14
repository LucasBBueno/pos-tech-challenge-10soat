package order

import (
	"context"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type ListOrders interface {
	Execute(ctx context.Context, limit uint64) (*domain.ListOrders, error)
}

type ListOrdersUseCase struct {
	orderRepository ports.OrderRepository
}

func CreateNewListOrdersUsecase(orderRepository ports.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		orderRepository,
	}
}

func (s *ListOrdersUseCase) Execute(ctx context.Context, limit uint64) (*domain.ListOrders, error) {
	orders, err := s.orderRepository.ListOrders(ctx, limit)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
