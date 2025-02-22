package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/r3iwan/mse-business-go/internal/models"
)

type CompanyRepository interface {
	CreateCompany(company *models.Companies) error
	IsEmailTaken(email string) (bool, error)
	GetCompanyPassword(email string) (string, error)
	ListCompanies() ([]*models.Companies, error)
	UpdateCompanyInfo(company *models.Companies) error
	DeleteCompany(id int) error
}

type companyRepository struct {
	conn *pgx.Conn
}

func NewCompanyRepository(conn *pgx.Conn) CompanyRepository {
	return &companyRepository{conn: conn}
}

func (r *companyRepository) CreateCompany(company *models.Companies) error {
	query := `
		INSERT INTO companies (company_name, company_address, company_phone, company_email, company_website, password)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING company_id
	`
	err := r.conn.QueryRow(context.Background(), query,
		company.CompanyName,
		company.CompanyAddress,
		company.CompanyPhone,
		company.CompanyEmail,
		company.CompanyWebsite,
		company.CompanyPassword,
	).Scan(&company.CompanyID)
	if err != nil {
		return fmt.Errorf("error inserting company: %w", err)
	}
	return nil
}

func (r *companyRepository) IsEmailTaken(email string) (bool, error) {
	query := `SELECT company_email FROM companies WHERE company_email = $1`
	var foundEmail string
	err := r.conn.QueryRow(context.Background(), query, email).Scan(&foundEmail)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *companyRepository) GetCompanyPassword(email string) (string, error) {
	query := `SELECT password FROM companies WHERE company_email = $1`

	var hashedPwd string
	err := r.conn.QueryRow(context.Background(), query, email).Scan(&hashedPwd)
	if err != nil {
		return "", err
	}

	return hashedPwd, err
}

func (r *companyRepository) ListCompanies() ([]*models.Companies, error) {
	query := `SELECT company_id, company_name, company_address, company_phone, company_email, company_website, password FROM companies`

	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []*models.Companies
	for rows.Next() {
		var company models.Companies
		err := rows.Scan(
			&company.CompanyID,
			&company.CompanyName,
			&company.CompanyAddress,
			&company.CompanyPhone,
			&company.CompanyEmail,
			&company.CompanyWebsite,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, &company)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return companies, nil
}

func (r *companyRepository) UpdateCompanyInfo(company *models.Companies) error {
	query := `UPDATE companies SET company_name = $1, company_address = $2, company_phone = $3, company_email = $4, company_website = $5, password = $6,
	WHERE company_id = $7`

	err := r.conn.QueryRow(context.Background(), query, company.CompanyName, company.CompanyAddress, company.CompanyPhone, company.CompanyEmail, company.CompanyWebsite, company.CompanyPassword, company.CompanyID).Scan(&company.CompanyID)

	return err
}

func (r *companyRepository) DeleteCompany(id int) error {
	query := `DELETE FROM companies WHERE company_id = $1`
	_, err := r.conn.Exec(context.Background(), query, id)
	return err
}
