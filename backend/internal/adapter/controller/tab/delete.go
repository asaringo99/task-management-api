package controller

import (
	"fmt"
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	condition_delete "github.com/asaringo99/task_management/internal/application/usecase/tab/delete/condition"
	"github.com/labstack/echo/v4"
)

func (controller TabController) Delete(c echo.Context) error {
	userid := authjwt.RetrieveUserIdFromToken(c)
	tabid, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(tabid, userid)
	condition := condition_delete.NewTabDeleteCondition(tabid, userid)
	err := controller.tabDeleteUsecase.Delete(condition)
	if err != nil {
		return err
	}
	return nil
}
