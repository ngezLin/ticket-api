package controller

import (
	"net/http"
	"ticket-api/service"

	"github.com/gin-gonic/gin"
)

type AdminReportController struct {
	TicketService service.TicketService
}

func NewAdminReportController(s service.TicketService) *AdminReportController {
	return &AdminReportController{s}
}

func (rc *AdminReportController) GetSalesReport(c *gin.Context) {
	report, err := rc.TicketService.GetSalesReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get sales report"})
		return
	}
	c.JSON(http.StatusOK, report)
}