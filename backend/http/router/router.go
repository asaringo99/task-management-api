package router

import (
	"github.com/asaringo99/task_management/http/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(handler []handler.Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	for _, h := range handler {
		h.AddEntryPoint(e)
	}
	return e
}
