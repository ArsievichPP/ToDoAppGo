package handlers

import (
	"ToDoAppGo/app/internal/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TodoHandler struct {
	service *services.TodoService
}

func MakeTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

func (td TodoHandler) GetTodoList(e echo.Context) error {
	list, err := td.service.GetList()
	if err != nil {
		return e.String(http.StatusOK, "")
	}

	return e.String(http.StatusOK, fmt.Sprintf("%+v", list))
}

func (td TodoHandler) GetToDoItem(c echo.Context) error {
	id := c.Get("id")
	todo, err := td.service.GetById(id.(uint64))
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}
	return c.String(http.StatusOK, fmt.Sprintf("%+v", todo))
}

func (td TodoHandler) CreateTodoItem(c echo.Context) error {
	return c.String(http.StatusOK, "ToDo create item")
}

func (td TodoHandler) UpdateTodoItem(c echo.Context) error {
	return c.String(http.StatusOK, "ToDo update item")
}

func (td TodoHandler) DeleteTodoItem(c echo.Context) error {
	return c.String(http.StatusOK, "ToDo delete item")
}
