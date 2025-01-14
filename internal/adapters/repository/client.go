package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"post-tech-challenge-10soat/internal/application/core/domain"
	"post-tech-challenge-10soat/internal/infra/storage/postgres"
)

type ClientRepository struct {
	db *postgres.DB
}

func NewClientRepository(db *postgres.DB) *ClientRepository {
	return &ClientRepository{
		db,
	}
}

func (repository *ClientRepository) CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	query := repository.db.QueryBuilder.Insert("clients").
		Columns("name", "email").
		Values(client.Name, client.Email).
		Suffix("RETURNING id, name, email, created_at, updated_at")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&client.Id,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (repository *ClientRepository) GetClientByCpf(ctx context.Context, cpf string) (*domain.Client, error) {
	var client domain.Client
	query := repository.db.QueryBuilder.Select("*").
		From("clients").
		Where(sq.Eq{"cpf": cpf}).
		Limit(1)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&client.Id,
		&client.Cpf,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (repository *ClientRepository) GetClientById(ctx context.Context, id string) (*domain.Client, error) {
	var client domain.Client
	query := repository.db.QueryBuilder.Select("*").
		From("clients").
		Where(sq.Eq{"id": id}).
		Limit(1)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&client.Id,
		&client.Cpf,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
