package domain

import (
	"time"
)

type OrderProduct struct {
	Id          string
	OrderId     string
	ProductId   string
	Quantity    int
	SubTotal    float64
	Observation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
