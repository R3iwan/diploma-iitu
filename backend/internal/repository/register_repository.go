package repository

import (
	"fmt"

	"github.com/r3iwan/mse-business-go/internal/models"
)

func RegisterCustomerToDB() error {
	var req models.RegisterCustomer

	if req.Email == "" || req.Username == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		return fmt.Errorf("Please fill all fields")
	}

	query := "INSERT INTO customers (username, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(query, req.Username, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		return fmt.Errorf("Error while inserting data to database: %v", err)
	}
}
