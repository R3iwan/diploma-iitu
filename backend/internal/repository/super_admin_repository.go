package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/r3iwan/mse-business-go/internal/models"
)

type SuperAdminRepository interface {
	CreateAdmin(admin *models.Admin) error
}

type superAdminRepository struct {
	conn *pgx.Conn
}

func NewSuperAdminRepository(conn *pgx.Conn) SuperAdminRepository {
	return &superAdminRepository{conn: conn}
}

func (r *superAdminRepository) CreateAdmin(admin *models.Admin) error {
	query := `INSERT INTO admins(admin_id, username, password, first_name, last_name, email, role)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := r.conn.QueryRow(context.Background(), query, admin.AdminID, admin.Username, admin.Password, admin.FirstName, admin.LastName, admin.Email, admin.Role).Scan(&admin.AdminID)

	return err
}
