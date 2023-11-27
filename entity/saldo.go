package entity

import (
	"time"

	"gorm.io/gorm"
)

type saldo struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Amount    float64        `json:"amount"`
	Currency  string         `json:"currency"`
	Timestamp time.Time      `json:"timestamp"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func topupsaldo(amount float64, name, currency string) *saldo {
	return &saldo{
		Name:      name,
		Amount:    amount,
		Currency:  currency,
		Timestamp: time.Now(),
	}
}

func ceksaldo(amount float64, currency string) *saldo {
	return &saldo{
		Amount:    amount,
		Currency:  currency,
		Timestamp: time.Now(),
	}
}
