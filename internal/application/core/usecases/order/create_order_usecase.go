package order

import (
	"context"
	"fmt"
	domain "post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/core/usecases/payment"
	"post-tech-challenge-10soat/internal/application/ports"

	"github.com/google/uuid"
)

type CreateOrder interface {
	Execute(ctx context.Context, createOrder *domain.CreateOrder) (*domain.Order, error)
}

type CreateOrderUsecase struct {
	productRepository      ports.ProductRepository
	clientRepository       ports.ClientRepository
	orderRepository        ports.OrderRepository
	orderProductRepository ports.OrderProductRepository
	paymentCheckout        payment.PaymentCheckout
}

func NewCreateOrderUsecase(
	productRepository ports.ProductRepository,
	clientRepository ports.ClientRepository,
	orderRepository ports.OrderRepository,
	orderProductRepository ports.OrderProductRepository,
	paymentCheckout payment.PaymentCheckout) CreateOrder {
	return &CreateOrderUsecase{
		productRepository,
		clientRepository,
		orderRepository,
		orderProductRepository,
		paymentCheckout,
	}
}

func (s *CreateOrderUsecase) Execute(ctx context.Context, createOrder *domain.CreateOrder) (*domain.Order, error) {
	var totalValue float64
	for _, orderProduct := range createOrder.Products {
		product, err := s.productRepository.GetProductById(ctx, orderProduct.ProductId)
		if err != nil {
			if err == domain.ErrDataNotFound {
				return nil, err
			}
			return nil, fmt.Errorf("cannot create order because has invalid product - %s", err.Error())
		}
		subTotal := product.Value * float64(orderProduct.Quantity)
		totalValue += subTotal
	}
	paymentInfo := domain.CreatePayment{
		Provider: domain.PaymentProviderMp,
		Type:     domain.PaymentTypePixQRCode,
	}
	p, err := s.paymentCheckout.Execute(ctx, &paymentInfo)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create order because cannot complete checkout - %s", err.Error())
	}
	orderInfo := domain.Order{
		Status:    domain.OrderStatusReceived,
		PaymentId: p.Id,
		Total:     totalValue,
	}
	if createOrder.ClientId != "" && uuid.Validate(createOrder.ClientId) == nil {
		client, err := s.clientRepository.GetClientById(ctx, createOrder.ClientId)
		if err != nil {
			if err == domain.ErrDataNotFound {
				return nil, err
			}
			return nil, fmt.Errorf("cannot create order because has invalid client - %s", err.Error())
		}
		orderInfo.ClientId = client.Id
	} else {
		orderInfo.ClientId = ""
	}
	order, err := s.orderRepository.CreateOrder(ctx, &orderInfo)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create order - %s", err.Error())
	}
	for _, orderProduct := range createOrder.Products {
		product, err := s.productRepository.GetProductById(ctx, orderProduct.ProductId)
		if err != nil {
			if err == domain.ErrDataNotFound {
				return nil, err
			}
			return nil, fmt.Errorf("cannot create order because has invalid product - %s", err.Error())
		}
		subTotal := product.Value * float64(orderProduct.Quantity)
		orderProductInfo := domain.OrderProduct{
			OrderId:     order.Id,
			ProductId:   product.Id,
			Quantity:    orderProduct.Quantity,
			SubTotal:    subTotal,
			Observation: orderProduct.Observation,
		}
		_, err = s.orderProductRepository.CreateOrderProduct(ctx, &orderProductInfo)
		if err != nil {
			if err == domain.ErrDataNotFound {
				return nil, err
			}
			err := s.orderRepository.DeleteOrder(ctx, order.Id)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("cannot complete order - %s", err.Error())
		}
	}
	return order, nil
}
