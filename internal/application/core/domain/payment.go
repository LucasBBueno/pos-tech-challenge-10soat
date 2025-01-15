package domain

import (
	"time"
)

const (
	PaymentTypePixQRCode = "PIX-QRCODE"
)

const (
	PaymentProviderMp = "mercado-pago"
)

type Payment struct {
	Id        string
	Provider  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreatePayment struct {
	Provider string
	Type     string
}
