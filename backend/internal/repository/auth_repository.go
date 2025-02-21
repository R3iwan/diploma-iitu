package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/r3iwan/mse-business-go/internal/models"
)

type AuthRepository interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomerByUsernameOrEmail(identifier string) (*models.Customer, error)
	UpdateCustomerDetails(customer *models.Customer) error
	GetPassword(identifier string) (string, error)
}

type authRepository struct {
	conn *pgx.Conn
}

func NewAuthRepository(conn *pgx.Conn) AuthRepository {
	return &authRepository{conn: conn}
}

func (r *authRepository) CreateCustomer(customer *models.Customer) error {
	query := `
		INSERT INTO customers 
			(company_id, username, first_name, last_name, email, role, phone, address, password, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING customer_id
	`

	now := time.Now()
	err := r.conn.QueryRow(
		context.Background(),
		query,
		customer.CompanyId,
		customer.Username,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.Role,
		customer.Phone,
		customer.Address,
		customer.Password,
		now,
		now,
	).Scan(&customer.CustomerID)

	return err
}

func (r *authRepository) GetCustomerByUsernameOrEmail(identifier string) (*models.Customer, error) {
	query := `
		SELECT customer_id, company_id, username, first_name, last_name, email, role, phone, address, password, created_at, updated_at
		FROM customers
		WHERE username = $1 OR email = $1
		LIMIT 1
	`
	var customer models.Customer
	err := r.conn.QueryRow(context.Background(), query, identifier).Scan(
		&customer.CustomerID,
		&customer.CompanyId,
		&customer.Username,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.Role,
		&customer.Phone,
		&customer.Address,
		&customer.Password,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *authRepository) UpdateCustomerDetails(customer *models.Customer) error {
	query := `
        UPDATE customers 
        SET username = $1, 
            first_name = $2, 
            last_name = $3, 
            email = $4, 
            phone = $5, 
            address = $6, 
            updated_at = $7
        WHERE customer_id = $8
        RETURNING customer_id
    `

	now := time.Now()
	err := r.conn.QueryRow(context.Background(), query,
		customer.Username,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.Phone,
		customer.Address,
		now,
		customer.CustomerID,
	).Scan(&customer.CustomerID)
	return err

}

func (r *authRepository) GetPassword(identifier string) (string, error) {
	query := `SELECT password FROM customers WHERE username = $1 OR email = $1 LIMIT 1`

	var hashedPwd string
	err := r.conn.QueryRow(context.Background(), query, identifier).Scan(&hashedPwd)
	if err != nil {
		return "", err
	}
	return hashedPwd, nil
}
