package services

import (
	"fmt"

	"github.com/r3iwan/mse-business-go/internal/models"
)

func RegisterCustomer() error {
	var req models.RegisterCustomer

	if req.Email == "" || req.Username == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		return fmt.Errorf("Please fill all fields")
	}

	
}
