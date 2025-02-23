package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/r3iwan/mse-business-go/internal/models"
)

type AdminRepository interface {
	CreateManager(manager *models.Manager) error
	CreateEmployee(employee *models.Employee) error
}

type adminRepository struct {
	conn *pgx.Conn
}

func NewAdminRepository(conn *pgx.Conn) AdminRepository {
	return &adminRepository{conn: conn}
}

func (r *adminRepository) CreateManager(manager *models.Manager) error {
	query := `INSERT INTO managers(manager_id, username, password, first_name, last_name, email, role)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := r.conn.QueryRow(context.Background(), query, manager.ManagerID, manager.Username, manager.Password, manager.FirstName, manager.LastName, manager.Email, manager.Role).Scan(&manager.ManagerID)

	return err
}

func (r *adminRepository) CreateEmployee(employee *models.Employee) error {
	query := `INSERT INTO employees(employee_id, username, password, first_name, last_name, email, role)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := r.conn.QueryRow(context.Background(), query, employee.EmployeeID, employee.Username, employee.Password, employee.FirstName, employee.LastName, employee.Email, employee.Role).Scan(&employee.EmployeeID)

	return err
}
