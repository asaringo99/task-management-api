package handler

import (
	"net/http"

	auth_cookie "github.com/asaringo99/task_management/http/auth/cookie"
	auth_jwt "github.com/asaringo99/task_management/http/auth/jwt"
	res "github.com/asaringo99/task_management/http/response"
	"github.com/labstack/echo/v4"
)

func (t *BoardHandler) post(c echo.Context) error {
	var response res.ResponseBody

	ret, err := t.controller.Post(c)
	if err != nil {
		response.Status = res.MessageError
		response.Error = err.Error()
		return c.JSON(http.StatusUnauthorized, response)
	}
	userid := auth_jwt.RetrieveUserIdFromToken(c)
	token, _ := auth_jwt.CreateJwtToken(userid, false)
	cookie := auth_cookie.NewCookieForJwtToken(token)
	c.SetCookie(cookie)
	response.Status = res.MessageSuccess
	response.Data = ret
	return c.JSON(http.StatusOK, response)
}
