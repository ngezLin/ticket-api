package controller

import (
	"net/http"
	"ticket-api/src/service"
	"ticket-api/src/utils"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	TicketService service.TicketService
}

func NewTicketController(s service.TicketService) *TicketController {
	return &TicketController{s}
}

func (tc *TicketController) Purchase(c *gin.Context) {
	var req struct {
		EventID  uint `json:"event_id"`
		Quantity int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID, err := utils.ExtractUserIDFromJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	ticket, err := tc.TicketService.Purchase(userID, req.EventID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ticket)
}
