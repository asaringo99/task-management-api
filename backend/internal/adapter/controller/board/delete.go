package controller

import (
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	condition_delete "github.com/asaringo99/task_management/internal/application/usecase/board/delete/condition"
	"github.com/labstack/echo/v4"
)

func (controller BoardController) Delete(c echo.Context) error {
	userid := authjwt.RetrieveUserIdFromToken(c)
	boardid, _ := strconv.Atoi(c.Param("boardid"))
	board := condition_delete.NewBoardDeleteCondition(userid, boardid)
	err := controller.boardDeleteUsecase.Delete(board)
	if err != nil {
		return err
	}
	return nil
}
