package handler

import (
	"net/http"

	"github.com/asaringo99/task_management/http/auth/controller/login"
	"github.com/asaringo99/task_management/http/auth/controller/token"
	auth_cookie "github.com/asaringo99/task_management/http/auth/cookie"
	"github.com/asaringo99/task_management/http/auth/entity"
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
		userCredential := new(entity.UserCredential)
		if err := c.Bind(userCredential); err != nil {
			response.Status = err.Error()
			return c.JSON(http.StatusInternalServerError, response)
		}
		if err := h.loginController.Login(*userCredential); err != nil {
			response.Status = res.MessageError
			return c.JSON(http.StatusNotAcceptable, response)
		}
		token, err := h.tokenController.RetrieveToken(*userCredential)
		if err != nil {
			response.Status = res.MessageError
			return c.JSON(http.StatusInternalServerError, response)
		}
		cookie := auth_cookie.NewCookieForJwtToken(token)
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, res.ResponseBody{
			Data: LoginResponseData{Token: token},
		})
	})
}
