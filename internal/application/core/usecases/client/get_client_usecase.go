package client

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/application/ports"
)

type GetClientUseCase struct {
	repository ports.ClientRepository
}

func NewGetClientUseCase(repository ports.ClientRepository) *GetClientUseCase {
	return &GetClientUseCase{
		repository,
	}
}

func (s *GetClientUseCase) Execute(ctx context.Context, cpf string) (*domain.Client, error) {
	client, err := s.repository.GetClientByCpf(ctx, cpf)
	if err != nil {
		return nil, fmt.Errorf("failed to get client by cpf - %s", err.Error())
	}
	return client, nil
}
