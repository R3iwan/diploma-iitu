package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/models"
	"github.com/r3iwan/mse-business-go/internal/services"
)

type CompHandler struct {
	compService services.CompanyService
}

func NewCompHandler(compService services.CompanyService) *CompHandler {
	return &CompHandler{compService: compService}
}

func (h *CompHandler) RegisterCompanyHandler(c *gin.Context) {
	var req models.Companies

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	err := h.compService.RegisterCompany(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error registering company", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "company registered successfully"})
}

func (h *CompHandler) LoginCompanyHandler(c *gin.Context) {
	var req models.LoginCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	err := h.compService.LoginCompany(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error log on customer", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "customer logged in  successfully"})
}
