package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/models"
	"github.com/r3iwan/mse-business-go/internal/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterCustomerHandler(c *gin.Context) {
	var req models.RegisterCustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	err := h.authService.RegisterCustomer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error registering customer", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "customer registered successfully"})
}

func (h *AuthHandler) LoginCustomerHanlder(c *gin.Context) {
	var req models.LoginCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	err := h.authService.LoginCustomer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error log on customer", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "customer logged in  successfully"})
}
