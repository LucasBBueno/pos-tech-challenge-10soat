package service

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/core/domain"
	"post-tech-challenge-10soat/internal/core/port"

	"github.com/google/uuid"
)

type OrderService struct {
	productRepository      port.ProductRepository
	clientRepository       port.ClientRepository
	orderRepository        port.OrderRepository
	orderProductRepository port.OrderProductRepository
	paymentService         port.PaymentService
}

func NewOrderService(productRepository port.ProductRepository, clientRepository port.ClientRepository, orderRepository port.OrderRepository, orderProductRepository port.OrderProductRepository, paymentService port.PaymentService) *OrderService {
	return &OrderService{
		productRepository,
		clientRepository,
		orderRepository,
		orderProductRepository,
		paymentService,
	}
}

func (service *OrderService) CreateOrder(ctx context.Context, createOrder *domain.CreateOrder) (*domain.Order, error) {
	var totalValue float64
	for _, orderProduct := range createOrder.Products {
		product, err := service.productRepository.GetProductById(ctx, orderProduct.ProductId)
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
	payment, err := service.paymentService.Checkout(ctx, &paymentInfo)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create order because cannot complete checkout - %s", err.Error())
	}
	orderInfo := domain.Order{
		Status:    domain.OrderStatusReceived,
		PaymentId: payment.Id,
		Total:     totalValue,
	}
	if createOrder.ClientId != "" && uuid.Validate(createOrder.ClientId) == nil {
		client, err := service.clientRepository.GetClientById(ctx, createOrder.ClientId)
		if err != nil {
			if err == domain.ErrDataNotFound {
				return nil, err
			}
			return nil, fmt.Errorf("cannot create order because has invalid client - %s", err.Error())
		}
		orderInfo.ClientId = &client.Id
	} else {
		orderInfo.ClientId = nil
	}
	order, err := service.orderRepository.CreateOrder(ctx, &orderInfo)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("cannot create order - %s", err.Error())
	}
	for _, orderProduct := range createOrder.Products {
		product, err := service.productRepository.GetProductById(ctx, orderProduct.ProductId)
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
		_, err = service.orderProductRepository.CreateOrderProduct(ctx, &orderProductInfo)
		if err != nil {
			if err == domain.ErrDataNotFound {
				return nil, err
			}
			service.orderRepository.DeleteOrder(ctx, order.Id.String())
			return nil, fmt.Errorf("cannot complete order - %s", err.Error())
		}
	}
	return order, nil
}

func (service *OrderService) ListOrders(ctx context.Context, limit uint64) (*domain.ListOrders, error) {
	orders, err := service.orderRepository.ListOrders(ctx, limit)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
