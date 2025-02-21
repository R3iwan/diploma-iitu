package services

import (
	"fmt"
	"strings"

	"github.com/r3iwan/mse-business-go/internal/models"
	"github.com/r3iwan/mse-business-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterCustomer(models.RegisterCustomerRequest) error
	LoginCustomer(models.LoginCustomerRequest) error
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) *authService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) RegisterCustomer(req models.RegisterCustomerRequest) error {
	req.Email = strings.TrimSpace(req.Email)
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	req.FirstName = strings.TrimSpace(req.FirstName)
	req.LastName = strings.TrimSpace(req.LastName)

	if req.Email == "" || req.Username == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		return fmt.Errorf("please fill all fields")
	}

	_, err := s.authRepo.GetCustomerByUsernameOrEmail(req.Username)
	if err == nil {
		return fmt.Errorf("username is already taken")
	}

	_, err = s.authRepo.GetCustomerByUsernameOrEmail(req.Email)
	if err == nil {
		return fmt.Errorf("email is already registered")
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	customer := &models.Customer{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      "customer",
	}

	err = s.authRepo.CreateCustomer(customer)
	if err != nil {
		return fmt.Errorf("error creating customer: %w", err)
	}

	return nil
}

func (s *authService) LoginCustomer(req models.LoginCustomerRequest) error {
	req.UsernameOrEmail = strings.TrimSpace(req.UsernameOrEmail)
	req.Password = strings.TrimSpace(req.Password)

	if req.UsernameOrEmail == "" || req.Password == "" {
		return fmt.Errorf("please fill all fields")
	}

	_, err := s.authRepo.GetCustomerByUsernameOrEmail(req.UsernameOrEmail)
	if err != nil {
		return fmt.Errorf("email or username not found")
	}

	login_pwd := req.Password
	hashedPwd, err := s.authRepo.GetPassword(req.UsernameOrEmail)
	if err != nil {
		return fmt.Errorf("error retrieving password: %v", err)
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

func hashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bcryptPassword), nil
}

func checkHashPassword(hashedPwd string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
