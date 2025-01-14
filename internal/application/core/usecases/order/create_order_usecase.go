package order

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	domain2 "post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/core/usecases/payment"
	"post-tech-challenge-10soat/internal/application/ports"
)

type CreateOrder interface {
	Execute(ctx context.Context, createOrder *domain2.CreateOrder) (*domain2.Order, error)
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

func (s *CreateOrderUsecase) Execute(ctx context.Context, createOrder *domain2.CreateOrder) (*domain2.Order, error) {
	var totalValue float64
	for _, orderProduct := range createOrder.Products {
		product, err := s.productRepository.GetProductById(ctx, orderProduct.ProductId)
		if err != nil {
			if err == domain2.ErrDataNotFound {
				return nil, err
			}
			return nil, fmt.Errorf("cannot create order because has invalid product - %s", err.Error())
		}
		subTotal := product.Value * float64(orderProduct.Quantity)
		totalValue += subTotal
	}
	paymentInfo := domain2.CreatePayment{
		Provider: domain2.PaymentProviderMp,
		Type:     domain2.PaymentTypePixQRCode,
	}
	p, err := s.paymentCheckout.Execute(ctx, &paymentInfo)
	if err != nil {
		if err == domain2.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create order because cannot complete checkout - %s", err.Error())
	}
	orderInfo := domain2.Order{
		Status:    domain2.OrderStatusReceived,
		PaymentId: p.Id,
		Total:     totalValue,
	}
	if createOrder.ClientId != "" && uuid.Validate(createOrder.ClientId) == nil {
		client, err := s.clientRepository.GetClientById(ctx, createOrder.ClientId)
		if err != nil {
			if err == domain2.ErrDataNotFound {
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
		if err == domain2.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create order - %s", err.Error())
	}
	for _, orderProduct := range createOrder.Products {
		product, err := s.productRepository.GetProductById(ctx, orderProduct.ProductId)
		if err != nil {
			if err == domain2.ErrDataNotFound {
				return nil, err
			}
			return nil, fmt.Errorf("cannot create order because has invalid product - %s", err.Error())
		}
		subTotal := product.Value * float64(orderProduct.Quantity)
		orderProductInfo := domain2.OrderProduct{
			OrderId:     order.Id,
			ProductId:   product.Id,
			Quantity:    orderProduct.Quantity,
			SubTotal:    subTotal,
			Observation: orderProduct.Observation,
		}
		_, err = s.orderProductRepository.CreateOrderProduct(ctx, &orderProductInfo)
		if err != nil {
			if err == domain2.ErrDataNotFound {
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
