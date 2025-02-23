package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/delivery"
	"github.com/r3iwan/mse-business-go/internal/services"
)

func RegisterCustomerRoutes(r *gin.Engine, authHandler delivery.AuthHandler) {
	customer := r.Group("/customer")
	{
		customer.POST("/register", authHandler.RegisterCustomerHandler)
		customer.POST("/login", authHandler.LoginCustomerHanlder)
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
		superAdmin.POST("/admin/register", superAdminHandler.RegisterAdminHandler)
		superAdmin.POST("/customer/register", superAdminHandler.RegisterCustomerHandler)
		superAdmin.POST("/company/register", superAdminHandler.RegisterCompanyHandler)
		superAdmin.POST("/customer/login", superAdminHandler.LoginCustomerHanlder)
		superAdmin.POST("/company/login", superAdminHandler.LoginCompanyHandler)
	}
}

func RegisterAdminRoutes(r *gin.Engine, adminService services.AdminService, authHandler delivery.AuthHandler) {
	adminHandler := delivery.NewAdminHandler(adminService, &authHandler)

	admin := r.Group("/admin")
	{
		admin.POST("/manager/register", adminHandler.RegisterManagerHandler)
		admin.POST("/employee/register", adminHandler.RegisterEmployeeHandler)
		admin.POST("/customer/register", adminHandler.RegisterCustomerHandler)
	}
}
