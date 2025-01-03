package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID
	Name        string
	Description string
	Image       string
	Value       float64
	CategoryId  uuid.UUID
	Category    *Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
