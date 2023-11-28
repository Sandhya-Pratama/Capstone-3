package entity

import (
	"time"
)

type Order struct {
	Id        int64      `json:"id"`
	TicketID  int64      `json:"ticket_id"`
	Ticket    Ticket     `json:"ticket"`
	UserID    int64      `json:"user_id"`
	User      User       `json:"user"`
	Quantity  int64      `json:"quantity"`
	Total     int64      `json:"total"`
	OrderAt   time.Time  `json:"order_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	OrderBy   string     `json:"order_by"`
	UpdateBy  string     `json:"-"`
	DeleteBy  string     `json:"-"`
}

// membuat func NewOrder dengan memanggil tiketID, quantity, total, dan OrderAt
func NewOrder(ticketID, quantity, userID int64) *Order {
	return &Order{
		TicketID: ticketID,
		Quantity: quantity,
		UserID:   userID,
		OrderAt:  time.Now(),
	}
}
