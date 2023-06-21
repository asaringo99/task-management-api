package controller

import (
	"fmt"
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/patch/condition"
	"github.com/labstack/echo/v4"
)

type TabPatchRequest struct {
	Column string      `json:"column"`
	Value  interface{} `json:"value"`
}

func (controller TabController) Patch(c echo.Context) error {
	userid := authjwt.RetrieveUserIdFromToken(c)
	tabid, _ := strconv.Atoi(c.Param("id"))
	request := new(TabPatchRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	fmt.Println(request, userid)
	input := condition.NewTabPatchCondition(tabid, userid, condition.PatchData{Column: request.Column, Value: request.Value})
	err := controller.tabPatchUsecase.Patch(input)
	if err != nil {
		return err
	}
	return nil
}
