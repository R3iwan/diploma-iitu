package services

import (
	"fmt"
	"strings"

	"github.com/r3iwan/mse-business-go/internal/models"
	"github.com/r3iwan/mse-business-go/internal/repository"
)

type CompanyService interface {
	RegisterCompany(models.Companies) error
	LoginCompany(models.LoginCompanyRequest) error
}

type compService struct {
	compRepo repository.CompanyRepository
}

func NewCompanyService(compRepo repository.CompanyRepository) *compService {
	return &compService{compRepo: compRepo}
}

func (s *compService) RegisterCompany(req models.Companies) error {
	req.CompanyName = strings.TrimSpace(req.CompanyName)
	req.CompanyAddress = strings.TrimSpace(req.CompanyAddress)
	req.CompanyPhone = strings.TrimSpace(req.CompanyPhone)
	req.CompanyEmail = strings.TrimSpace(req.CompanyEmail)
	req.CompanyWebsite = strings.TrimSpace(req.CompanyWebsite)
	req.CompanyPassword = strings.TrimSpace(req.CompanyPassword)

	if req.CompanyName == "" || req.CompanyAddress == "" || req.CompanyPhone == "" || req.CompanyEmail == "" || req.CompanyWebsite == "" || req.CompanyPassword == "" {
		return fmt.Errorf("please fill all fields")
	}

	taken, err := s.compRepo.IsEmailTaken(req.CompanyEmail)
	if err != nil {
		return fmt.Errorf("unexpected error: %v", err)
	}
	if taken {
		return fmt.Errorf("email is already taken")
	}

	hashedPassword, err := hashPassword(req.CompanyPassword)
	if err != nil {
		return err
	}

	company := &models.Companies{
		CompanyName:     req.CompanyName,
		CompanyAddress:  req.CompanyAddress,
		CompanyPhone:    req.CompanyPhone,
		CompanyEmail:    req.CompanyEmail,
		CompanyWebsite:  req.CompanyWebsite,
		CompanyPassword: hashedPassword,
	}

	err = s.compRepo.CreateCompany(company)
	if err != nil {
		return fmt.Errorf("error creating new company")
	}

	return err
}

func (s *compService) LoginCompany(req models.LoginCompanyRequest) error {
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Password == "" {
		return fmt.Errorf("error logging into company page")
	}

	login_pwd := req.Password
	hashedPwd, err := s.compRepo.GetCompanyPassword(req.Email)
	if err != nil {
		return fmt.Errorf("incorrect password")
	}

	match, err := checkHashPassword(hashedPwd, login_pwd)
	if err != nil {
		return fmt.Errorf("error when comparing password: %v", err)
	}
	if !match {
		return fmt.Errorf("invalid credentials")
	}

	return nil
}
