package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/delivery"
	"github.com/r3iwan/mse-business-go/internal/services"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler delivery.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register/customer", authHandler.RegisterCustomerHandler)
		auth.POST("/login/customer", authHandler.LoginCustomerHanlder)
	}
}

func RegisterCompanyRoutes(r *gin.Engine, compHandler delivery.CompHandler) {
	company := r.Group("/company")
	{
		company.POST("/register", compHandler.RegisterCompanyHandler)
		company.POST("/login", compHandler.LoginCompanyHandler)
	}
}

func RegisterSuperAdminRoutes(r *gin.Engine, superAdminServices services.SuperAdminServices, authHandler delivery.AuthHandler, compHandler delivery.CompHandler) {
	superAdminHandler := delivery.NewSuperAdminHandler(superAdminServices, &authHandler, &compHandler)

	superAdmin := r.Group("/super_admin")
	{
		superAdmin.POST("/register/admin", superAdminHandler.RegisterAdminHandler)
		superAdmin.POST("/register/customer", superAdminHandler.RegisterCustomerHandler)
		superAdmin.POST("/register/company", superAdminHandler.RegisterCompanyHandler)
		superAdmin.POST("/login/customer", superAdminHandler.LoginCustomerHanlder)
		superAdmin.POST("/login/company", superAdminHandler.LoginCompanyHandler)
	}
}
