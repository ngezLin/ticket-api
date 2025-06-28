package controller

import (
	"net/http"
	"ticket-api/src/entity"
	"ticket-api/src/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{UserService: us}
}

func (uc *UserController) Register(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.UserService.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "registration success",
		"user": gin.H{
			"id":    createdUser.ID,
			"name":  createdUser.Name,
			"email": createdUser.Email,
			"role":  createdUser.Role,
		},
	})
}

func (uc *UserController) Login(c *gin.Context){
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, token, err := uc.UserService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

		c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}