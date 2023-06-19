package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *TaskHandler) delete(c echo.Context) error {
	err := t.controller.Delete(c)
	if err != nil {
		return c.String(http.StatusNoContent, "Error!")
	}
	return c.String(http.StatusOK, "Success")
}
