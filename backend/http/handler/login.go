package handler

import (
	"net/http"

	"github.com/asaringo99/task_management/http/auth/controller/login"
	"github.com/asaringo99/task_management/http/auth/controller/token"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	loginController login.LoginController
	tokenController token.TokenController
}

func NewLoginHandler(loginController *login.LoginController, tokenController *token.TokenController) *LoginHandler {
	return &LoginHandler{
		loginController: *loginController,
		tokenController: *tokenController,
	}
}

func (h *LoginHandler) AddEntryPoint(e *echo.Echo) {
	e.POST("/login", func(c echo.Context) error {
		if err := h.loginController.Login(c); err != nil {
			return c.String(http.StatusNotAcceptable, "Not Found!")
		}
		token, err := h.tokenController.RetrieveToken(c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Not Created Token")
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": token,
		})
	})
}
