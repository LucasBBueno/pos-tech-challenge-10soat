package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
