package models

type RegisterCustomerRequest struct {
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type LoginCompanyRequest struct {
	Email    string `json:"company_email" binding:"required"`
	Password string `json:"company_password" binding:"required"`
}

type LoginCustomerRequest struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}
