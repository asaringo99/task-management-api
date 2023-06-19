package handler

import (
	"fmt"
	"net/http"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	res "github.com/asaringo99/task_management/http/response"
	"github.com/labstack/echo/v4"
)

func (t *BoardHandler) post(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Value)
	var response res.ResponseBody

	ret, err := t.controller.Post(c)
	if err != nil {
		response.Status = res.MessageError
		response.Error = err.Error()
		return c.JSON(http.StatusUnauthorized, response)
	}
	userid := authjwt.RetrieveUserIdFromToken(c)
	token, _ := authjwt.CreateJwtToken(userid, false)
	response.Status = res.MessageSuccess
	response.Data = ret
	response.Token = token
	return c.JSON(http.StatusOK, response)
}
