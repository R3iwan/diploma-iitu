package models

import "time"

type Order struct {
	OrderID     int       `json:"order_id"`
	ProductID   int       `json:"product_id"`
	CustomerID  int       `json:"customer_id"`
	OrderDate   time.Time `json:"order_date"`
	OrderStatus string    `json:"order_status"`
	TotalPrice  float64   `json:"total_price"`
}
