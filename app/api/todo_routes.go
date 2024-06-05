package api

import (
	"ToDoAppGo/app/internal/handlers"
	"ToDoAppGo/app/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitTodoRoutes(e *echo.Echo, db *gorm.DB) {
	todoGroup := e.Group("/todo")
	todoHandler := handlers.MakeTodoHandler(services.NewTodoService(db))

	todoGroup.GET("", todoHandler.GetTodoList)
	todoGroup.POST("", todoHandler.CreateTodoItem)
	todoGroup.GET("/:id", todoHandler.GetToDoItem)
	todoGroup.PATCH("/:id", todoHandler.UpdateTodoItem)
	todoGroup.DELETE("/:id", todoHandler.DeleteTodoItem)
}
