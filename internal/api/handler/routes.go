package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zuu-development/fullstack-examination-2024/internal/domain/service"
	"github.com/zuu-development/fullstack-examination-2024/internal/infrastructure/datastores/sqlite"
	"gorm.io/gorm"
)

// Register registers the routes for the application.
func Register(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	api := e.Group("/api/v1")

	// Health check
	healthHandler := NewHealth()
	api.GET("/healthz", healthHandler.Healthz)

	// Todo
	repository := sqlite.NewTodo(db)
	service := service.NewTodo(repository)
	todoHandler := NewTodo(service)
	todo := api.Group("/todos")
	{
		todo.POST("", todoHandler.Create)
		todo.GET("", todoHandler.FindAll)
		todo.GET("/:id", todoHandler.Find)
		todo.PUT("/:id", todoHandler.Update)
		todo.DELETE("/:id", todoHandler.Delete)
	}
}
