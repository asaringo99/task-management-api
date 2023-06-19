package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *TaskHandler) put(c echo.Context) error {
	err := t.controller.Put(c)
	if err != nil {
		return c.String(http.StatusNoContent, "Error!")
	}
	return c.String(http.StatusOK, "Success")
}
