package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/delivery"
	"github.com/r3iwan/mse-business-go/internal/services"
)

func RegisterAuthRoutes(r *gin.Engine, authService services.AuthService, compService services.CompanyService) {
	customer_handler := delivery.NewAuthHandler(authService)
	company_handler := delivery.NewCompHandler(compService)

	auth := r.Group("/auth")
	{
		auth.POST("/register/customer", customer_handler.RegisterCustomerHandler)
		auth.POST("/login/customer", customer_handler.LoginCustomerHanlder)
		auth.POST("/register/company", company_handler.RegisterCompanyHandler)
		auth.POST("/login/company", company_handler.LoginCompanyHandler)
	}

}
