package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderProduct struct {
	Id          uuid.UUID
	OrderId     uuid.UUID
	ProductId   uuid.UUID
	Quantity    int
	SubTotal    float64
	Observation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
