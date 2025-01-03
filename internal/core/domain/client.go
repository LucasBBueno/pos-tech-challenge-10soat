package domain

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id        uuid.UUID
	Cpf       *string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
