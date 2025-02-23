package services

import (
	"github.com/r3iwan/mse-business-go/internal/models"
	"github.com/r3iwan/mse-business-go/internal/repository"
)

type SuperAdminServices interface {
	RegisterAdmin(models.Admin) error
	RegisterCustomer(models.RegisterCustomerRequest) error
	LoginCustomer(models.LoginCustomerRequest) error
	RegisterCompany(models.Companies) error
	LoginCompany(models.LoginCompanyRequest) error
}

type superAdminServices struct {
	superAdminRepo repository.SuperAdminRepository
	authService    AuthService
	compService    CompanyService
}

func NewSuperAdminServices(superAdminRepo repository.SuperAdminRepository, authService AuthService, compService CompanyService) SuperAdminServices {
	return &superAdminServices{
		superAdminRepo: superAdminRepo,
		authService:    authService,
		compService:    compService,
	}
}

func (s *superAdminServices) RegisterAdmin(admin models.Admin) error {
	return s.superAdminRepo.CreateAdmin(&admin)
}

func (s *superAdminServices) RegisterCustomer(req models.RegisterCustomerRequest) error {
	return s.authService.RegisterCustomer(req)
}

func (s *superAdminServices) LoginCustomer(req models.LoginCustomerRequest) error {
	return s.authService.LoginCustomer(req)
}

func (s *superAdminServices) RegisterCompany(req models.Companies) error {
	return s.compService.RegisterCompany(req)
}

func (s *superAdminServices) LoginCompany(req models.LoginCompanyRequest) error {
	return s.compService.LoginCompany(req)
}
