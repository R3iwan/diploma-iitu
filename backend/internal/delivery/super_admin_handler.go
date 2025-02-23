package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/models"
	"github.com/r3iwan/mse-business-go/internal/services"
)

type SuperAdminHandler struct {
	superAdminService services.SuperAdminServices
	AuthHandler       *AuthHandler
	CompHandler       *CompHandler
}

func NewSuperAdminHandler(superAdminService services.SuperAdminServices, AuthHandler *AuthHandler, CompHandler *CompHandler) *SuperAdminHandler {
	return &SuperAdminHandler{superAdminService: superAdminService, AuthHandler: AuthHandler, CompHandler: CompHandler}
}

func (h *SuperAdminHandler) RegisterAdminHandler(c *gin.Context) {
	var req models.Admin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	err := h.superAdminService.RegisterAdmin(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error registering admin", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "admin registered successfully"})
}

func (h *SuperAdminHandler) RegisterCustomerHandler(c *gin.Context) {
	h.AuthHandler.RegisterCustomerHandler(c)
}

func (h *SuperAdminHandler) LoginCustomerHanlder(c *gin.Context) {
	h.AuthHandler.LoginCustomerHanlder(c)
}

func (h *SuperAdminHandler) RegisterCompanyHandler(c *gin.Context) {
	h.CompHandler.RegisterCompanyHandler(c)
}

func (h *SuperAdminHandler) LoginCompanyHandler(c *gin.Context) {
	h.CompHandler.LoginCompanyHandler(c)
}
