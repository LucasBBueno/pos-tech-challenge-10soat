package service

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/core/domain"
	"post-tech-challenge-10soat/internal/core/port"
)

type ClientService struct {
	repository port.ClientRepository
}

func NewClientService(repository port.ClientRepository) *ClientService {
	return &ClientService{
		repository,
	}
}

func (service *ClientService) CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	client, err := service.repository.CreateClient(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create client - %s", err.Error())
	}
	return client, nil
}

func (service *ClientService) GetClientByCpf(ctx context.Context, cpf string) (*domain.Client, error) {
	client, err := service.repository.GetClientByCpf(ctx, cpf)
	if err != nil {
		return nil, fmt.Errorf("failed to get client by cpf - %s", err.Error())
	}
	return client, nil
}
