package models

import "time"

type Customer struct {
	CompanyId  int       `json:"company_id"`
	CustomerID int       `json:"customer_id"`
	Username   string    `json:"username"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
