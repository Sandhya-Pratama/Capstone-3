package service

import (
	"Ticketing/entity"
	"context"
)

// cek saldo
type Ceksaldo interface {
	Ceksaldo(ctx context.Context, amount float64, name, currency string) (*entity.Saldo, error)
}
