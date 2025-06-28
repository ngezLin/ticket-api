package router

import (
	"ticket-api/src/controller"
	"ticket-api/src/middleware"
	"ticket-api/src/repository"
	"ticket-api/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Repositories
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	eventRepo := repository.NewEventRepository(db)
	ticketRepo := repository.NewTicketRepository(db)

	// Services
	userService := service.NewUserService(userRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	eventService := service.NewEventService(eventRepo)
	ticketService := service.NewTicketService(ticketRepo, eventRepo, userRepo)

	// Controllers
	userController := controller.NewUserController(userService)
	categoryController := controller.NewCategoryController(categoryService)
	eventController := controller.NewEventController(eventService)
	customerEventController := controller.NewCustomerEventController(eventService)
	customerBalanceController := controller.NewCustomerBalanceController(userService)
	ticketController := controller.NewTicketController(ticketService)
	customerTicketController := controller.NewCustomerTicketController(ticketService)
	ReportController := controller.NewAdminReportController(ticketService)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	//admin routes
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.RoleMiddleware("admin"))
	{
		adminGroup.GET("/categories", categoryController.GetAll)
		adminGroup.POST("/categories", categoryController.Create)
		adminGroup.PUT("/categories/:id", categoryController.Update)
		adminGroup.DELETE("/categories/:id", categoryController.Delete)

		adminGroup.GET("/events", eventController.GetAll)
		adminGroup.POST("/events", eventController.Create)
		adminGroup.PUT("/events/:id", eventController.Update)
		adminGroup.DELETE("/events/:id", eventController.Delete)
		adminGroup.GET("/reports/sales", ReportController.GetSalesReport)
	}

	// Customer routes
	customerGroup := r.Group("/customer")
	customerGroup.Use(middleware.RoleMiddleware("customer"))
	{
		customerGroup.GET("/events", customerEventController.GetActiveEvents)
		customerGroup.PUT("/me/balance", customerBalanceController.UpdateBalance)
		customerGroup.POST("/tickets", ticketController.Purchase)
		customerGroup.GET("/tickets", customerTicketController.GetMyTickets)
		customerGroup.GET("/tickets/:id",  customerTicketController.GetMyTicket)
		customerGroup.PATCH("/tickets/:id/cancel", customerTicketController.Cancel)

	}
}
