package router

import (
	"ticket-api/controller"
	"ticket-api/middleware"
	"ticket-api/repository"
	"ticket-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.RoleMiddleware("admin")) // FIXED HERE
	{
		adminGroup.GET("/categories", categoryController.GetAll)
		adminGroup.POST("/categories", categoryController.Create)
		adminGroup.PUT("/categories/:id", categoryController.Update)
		adminGroup.DELETE("/categories/:id", categoryController.Delete)

		adminGroup.GET("/events", eventController.GetAll)
		adminGroup.POST("/events", eventController.Create)
		adminGroup.PUT("/events/:id", eventController.Update)
		adminGroup.DELETE("/events/:id", eventController.Delete)
	}
}
