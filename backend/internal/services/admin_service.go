package services

import (
	"github.com/r3iwan/mse-business-go/internal/models"
	"github.com/r3iwan/mse-business-go/internal/repository"
)

type AdminService interface {
	CreateManager(models.Manager) error
	CreateEmployee(models.Employee) error
	CreateCustomer(models.RegisterCustomerRequest) error
}

type adminService struct {
	adminRepo   repository.AdminRepository
	authService AuthService
}

func NewAdminService(adminRepo repository.AdminRepository, authService AuthService) AdminService {
	return &adminService{adminRepo: adminRepo, authService: authService}
}

func (s *adminService) CreateManager(req models.Manager) error {
	return s.adminRepo.CreateManager(&req)
}

func (s *adminService) CreateEmployee(req models.Employee) error {
	return s.adminRepo.CreateEmployee(&req)
}

func (s *adminService) CreateCustomer(req models.RegisterCustomerRequest) error {
	return s.authService.RegisterCustomer(req)
}
