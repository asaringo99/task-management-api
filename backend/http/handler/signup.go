package handler

import (
	"net/http"

	"github.com/asaringo99/task_management/http/auth/controller/signup"
	res "github.com/asaringo99/task_management/http/response"
	"github.com/labstack/echo/v4"
)

type SignUpHandler struct {
	controller signup.SignUpController
}

func NewSignUpHandler(controller *signup.SignUpController) *SignUpHandler {
	return &SignUpHandler{
		controller: *controller,
	}
}

func (h *SignUpHandler) AddEntryPoint(e *echo.Echo) {
	var response res.ResponseBody

	e.POST("/signup", func(c echo.Context) error {
		if err := h.controller.SignUp(c); err != nil {
			response.Status = res.MessageError
			response.Error = err.Error()
			return c.JSON(http.StatusConflict, response)
		}
		return c.String(http.StatusOK, "CREATE Success!")
	})
}
