package handler

import (
	"net/http"

	"github.com/asaringo99/task_management/http/auth/controller/signup"
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
	e.POST("/signup", func(c echo.Context) error {
		h.controller.SignUp(c)
		return c.String(http.StatusOK, "CREATE Success!")
	})
}
