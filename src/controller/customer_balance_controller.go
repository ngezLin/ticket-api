package controller

import (
	"net/http"
	"ticket-api/src/service"

	"ticket-api/src/utils"

	"github.com/gin-gonic/gin"
)

type CustomerBalanceController struct {
	UserService service.UserService
}

func NewCustomerBalanceController(s service.UserService) *CustomerBalanceController {
	return &CustomerBalanceController{s}
}

func (cb *CustomerBalanceController) UpdateBalance(c *gin.Context) {
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
		return
	}

	userID, err := utils.ExtractUserIDFromJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	updated, err := cb.UserService.UpdateBalance(userID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
