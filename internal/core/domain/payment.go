package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	PaymentTypePixQRCode = "PIX-QRCODE"
)

const (
	PaymentProviderMp = "mercado-pago"
)

type Payment struct {
	Id        uuid.UUID
	Provider  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreatePayment struct {
	Provider string
	Type     string
}
