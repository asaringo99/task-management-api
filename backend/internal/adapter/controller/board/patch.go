package controller

import (
	"fmt"
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	"github.com/asaringo99/task_management/internal/application/usecase/board/patch/condition"
	"github.com/labstack/echo/v4"
)

type BoardPatchRequest struct {
	Column string      `json:"column"`
	Value  interface{} `json:"value"`
}

func (controller BoardController) Patch(c echo.Context) error {
	userid := authjwt.RetrieveUserIdFromToken(c)
	boardid, _ := strconv.Atoi(c.Param("id"))
	request := new(BoardPatchRequest)
	fmt.Println(userid)
	if err := c.Bind(request); err != nil {
		return err
	}
	fmt.Println(request, userid)
	input := condition.NewBoardPatchCondition(boardid, userid, condition.PatchData{Column: request.Column, Value: request.Value})
	err := controller.boardPatchUsecase.Patch(input)
	if err != nil {
		return err
	}
	return nil
}
