package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/asaringo99/task_management/http/auth/controller/login"
	"github.com/asaringo99/task_management/http/auth/controller/token"
	res "github.com/asaringo99/task_management/http/response"
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

type LoginResponseData struct {
	Token string `json:"token"`
}

func (h *LoginHandler) AddEntryPoint(e *echo.Echo) {
	var response res.ResponseBody

	e.POST("/login", func(c echo.Context) error {
		if err := h.loginController.Login(c); err != nil {
			response.Status = res.MessageError
			return c.JSON(http.StatusNotAcceptable, response)
		}
		token, err := h.tokenController.RetrieveToken(c)
		fmt.Println(token)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Not Created Token")
		}

		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = token
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.HttpOnly = true
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, res.ResponseBody{
			Data: LoginResponseData{Token: token},
		})
	})
}
