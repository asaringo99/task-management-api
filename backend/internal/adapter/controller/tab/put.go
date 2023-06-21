package controller

import (
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	condition_put "github.com/asaringo99/task_management/internal/application/usecase/tab/put/condition"
	"github.com/labstack/echo/v4"
)

func (controller TabController) Put(c echo.Context) error {
	// TODO: Bindで受け取る
	userid := authjwt.RetrieveUserIdFromToken(c)
	tabid, _ := strconv.Atoi(c.FormValue("tabid"))
	isActive, _ := strconv.ParseBool(c.FormValue("isactive"))
	title := c.FormValue("priority")
	input := condition_put.NewTabPutCondition(tabid, userid, isActive, title)
	err := controller.tabPutUsecase.Put(input)
	if err != nil {
		return err
	}
	return nil
}
