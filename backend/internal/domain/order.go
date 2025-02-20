package domain

import "time"

type Order struct {
	ID          int
	OrderNumber string
	TotalAmount float64
	Status      string
	CreatedAt   time.Time
}
