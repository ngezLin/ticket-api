package controller

import (
	"net/http"
	"strconv"
	"ticket-api/service"
	"ticket-api/utils"

	"github.com/gin-gonic/gin"
)

type CustomerTicketController struct {
	TicketService service.TicketService
}

func NewCustomerTicketController(s service.TicketService) *CustomerTicketController {
	return &CustomerTicketController{s}
}

func (ct *CustomerTicketController) GetMyTickets(c *gin.Context) {
	userID, err := utils.ExtractUserIDFromJWT(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	tickets, err := ct.TicketService.GetMyTickets(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H
	for _, ticket := range tickets {
		event := ticket.Event
		response = append(response, gin.H{
			"ticket_id": ticket.ID,
			"quantity":  ticket.Quantity,
			"status":    ticket.Status,
			"event": gin.H{
				"id":          event.ID,
				"name":        event.Name,
				"description": event.Description,
				"category":    event.Category.Name,
				"price":       event.Price,
				"start_time":  event.StartTime,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

func (ct *CustomerTicketController) GetMyTicket(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    userID, _ := utils.ExtractUserIDFromJWT(c)

    ticket, err := ct.TicketService.GetMyTicketByID(userID, uint(id))
    if err != nil {
        status := http.StatusBadRequest
        if err.Error() == "forbidden" { status = http.StatusForbidden }
        c.JSON(status, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, ticket)
}

func (ct *CustomerTicketController) Cancel(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    userID, _ := utils.ExtractUserIDFromJWT(c)

    ticket, err := ct.TicketService.Cancel(userID, uint(id))
    if err != nil {
        status := http.StatusBadRequest
        if err.Error() == "forbidden" { status = http.StatusForbidden }
        c.JSON(status, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, ticket)
}
