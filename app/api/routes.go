package api

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is OK!!!")
	})
	InitTodoRoutes(e, db)
}
