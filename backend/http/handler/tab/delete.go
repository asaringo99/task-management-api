package handler

import (
	"net/http"

	res "github.com/asaringo99/task_management/http/response"
	"github.com/labstack/echo/v4"
)

func (t *TabHandler) delete(c echo.Context) error {
	var response res.ResponseBody
	err := t.controller.Delete(c)
	if err != nil {
		response.Status = res.MessageError
		response.Error = err.Error()
		return c.JSON(http.StatusNotFound, response)
	}
	response.Status = res.MessageSuccess
	return c.JSON(http.StatusOK, response)

}
