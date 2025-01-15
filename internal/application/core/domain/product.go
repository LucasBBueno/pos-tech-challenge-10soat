package domain

import (
	"time"
)

type Product struct {
	Id          string
	Name        string
	Description string
	Image       string
	Value       float64
	CategoryId  string
	Category    *Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
