package controller

import (
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	condition_put "github.com/asaringo99/task_management/internal/application/usecase/board/put/condition"
	"github.com/labstack/echo/v4"
)

func (controller BoardController) Put(c echo.Context) error {
	// TODO: Bindで受け取る
	userid := authjwt.RetrieveUserIdFromToken(c)
	boardid, _ := strconv.Atoi(c.FormValue("boardid"))
	status := c.FormValue("status")
	priority, _ := strconv.Atoi(c.FormValue("priority"))
	input := condition_put.NewBoardPutCondition(boardid, userid, priority, status)
	err := controller.boardPutUsecase.Put(input)
	if err != nil {
		return err
	}
	return nil
}
