package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/models"
	"github.com/r3iwan/mse-business-go/internal/services"
)

type AdminHandler struct {
	adminService services.AdminService
	AuthHandler  *AuthHandler
}

func NewAdminHandler(adminService services.AdminService, AuthHandler *AuthHandler) *AdminHandler {
	return &AdminHandler{adminService: adminService, AuthHandler: AuthHandler}
}

func (h *AdminHandler) RegisterManagerHandler(c *gin.Context) {
	var req models.Manager
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	err := h.adminService.CreateManager(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error registering manager", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "manager registered successfully"})
}

func (h *AdminHandler) RegisterEmployeeHandler(c *gin.Context) {
	var req models.Employee
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	err := h.adminService.CreateEmployee(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error registering employee", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee registered successfully"})
}

func (h *AdminHandler) RegisterCustomerHandler(c *gin.Context) {
	h.AuthHandler.RegisterCustomerHandler(c)
}
