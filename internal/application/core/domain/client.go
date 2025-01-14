package domain

import (
	"time"
)

type Client struct {
	Id        string
	Cpf       string // TODO - pode se alterar para um objeto de valor (valueObject) inclusive outros campos
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
