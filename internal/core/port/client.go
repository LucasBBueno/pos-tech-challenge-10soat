package port

import (
	"context"
	"post-tech-challenge-10soat/internal/core/domain"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error)
	GetClientByCpf(ctx context.Context, cpf string) (*domain.Client, error)
}

type ClientService interface {
	CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error)
	GetClientByCpf(ctx context.Context, cpf string) (*domain.Client, error)
}
