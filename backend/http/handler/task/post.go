package handler

import (
	"net/http"

	res "github.com/asaringo99/task_management/http/response"
	"github.com/labstack/echo/v4"
)

func (t *TaskHandler) post(c echo.Context) error {
	var response res.ResponseBody
	ret, err := t.controller.Post(c)
	if err != nil {
		response.Status = res.MessageError
		response.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}
	response.Status = res.MessageSuccess
	response.Data = ret
	return c.JSON(http.StatusOK, response)
}
