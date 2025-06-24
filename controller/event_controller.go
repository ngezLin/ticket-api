package controller

import (
	"net/http"
	"strconv"
	"ticket-api/entity"
	"ticket-api/service"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	EventService service.EventService
}

func NewEventController(s service.EventService) *EventController {
	return &EventController{s}
}

func (ec *EventController) GetAll(c *gin.Context) {
	events, _ := ec.EventService.GetAll()
	c.JSON(http.StatusOK, events)
}

func (ec *EventController) Create(c *gin.Context) {
	var input entity.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := ec.EventService.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (ec *EventController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input entity.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := ec.EventService.Update(uint(id), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (ec *EventController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ec.EventService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
