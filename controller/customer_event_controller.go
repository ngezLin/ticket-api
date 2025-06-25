package controller

import (
	"net/http"
	"ticket-api/service"

	"github.com/gin-gonic/gin"
)

type CustomerEventController struct {
	EventService service.EventService
}

func NewCustomerEventController(s service.EventService) *CustomerEventController {
	return &CustomerEventController{s}
}

func (c *CustomerEventController) GetActiveEvents(ctx *gin.Context) {
	events, err := c.EventService.GetAllActive()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get events"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}
