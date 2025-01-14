package client

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type CreateClientUseCase struct {
	repository ports.ClientRepository
}

func NewCreateClientUsecase(repository ports.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{
		repository,
	}
}

func (s *CreateClientUseCase) Execute(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	client, err := s.repository.CreateClient(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create client - %s", err.Error())
	}
	return client, nil
}
