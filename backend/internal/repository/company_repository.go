package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/r3iwan/mse-business-go/internal/models"
)

type CompanyRepository interface {
	CreateCompany(company *models.Companies) error
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
		INSERT INTO companies (company_name, company_address, company_phone, company_email, company_website)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING company_id
	`
	err := r.conn.QueryRow(context.Background(), query,
		company.CompanyName,
		company.CompanyAddress,
		company.CompanyPhone,
		company.CompanyEmail,
		company.CompanyWebsite,
	).Scan(&company.CompanyID)
	if err != nil {
		return fmt.Errorf("error inserting company: %w", err)
	}
	return nil
}

func (r *companyRepository) ListCompanies() ([]*models.Companies, error) {
	query := `SELECT company_id, company_name, company_address, company_phone, company_email, company_website FROM companies`

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
	query := `UPDATE companies SET company_name = $1, company_address = $2, company_phone = $3, company_email = $4, company_website = $5
	WHERE company_id = $6`

	err := r.conn.QueryRow(context.Background(), query, company.CompanyName, company.CompanyAddress, company.CompanyPhone, company.CompanyEmail, company.CompanyWebsite, company.CompanyID).Scan(&company.CompanyID)

	return err
}

func (r *companyRepository) DeleteCompany(id int) error {
	query := `DELETE FROM companies WHERE company_id = $1`
	_, err := r.conn.Exec(context.Background(), query, id)
	return err
}
