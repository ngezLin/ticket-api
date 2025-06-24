package controller

import (
	"net/http"
	"strconv"
	"ticket-api/entity"
	"ticket-api/service"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(cs service.CategoryService) *CategoryController {
	return &CategoryController{cs}
}

func (cc *CategoryController) GetAll(c *gin.Context) {
	categories, _ := cc.CategoryService.GetAll()
	c.JSON(http.StatusOK, categories)
}

func (cc *CategoryController) Create(c *gin.Context) {
	var input entity.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := cc.CategoryService.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (cc *CategoryController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input entity.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := cc.CategoryService.Update(uint(id), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (cc *CategoryController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := cc.CategoryService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
