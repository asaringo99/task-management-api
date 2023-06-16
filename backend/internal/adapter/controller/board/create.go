package controller

import (
	"fmt"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	presenter_create "github.com/asaringo99/task_management/internal/adapter/presenter/board/create"
	usecase_create "github.com/asaringo99/task_management/internal/application/usecase/board/create"
	"github.com/labstack/echo/v4"
)

type Board struct {
	Status   string `json:"status"`
	Priority int    `json:"priority"`
}

func (controller BoardController) Post(c echo.Context) (*presenter_create.BoardCreatePresenterOutputDto, error) {
	// TODO: Bindで受け取る
	userid := authjwt.RetrieveUserIdFromToken(c)
	board := new(Board)
	if err := c.Bind(board); err != nil {
		return nil, err
	}
	fmt.Println(userid, board.Status, board.Priority)
	boardCureateUsecaseInput := usecase_create.NewBoardCreateUsecaseInput(userid, board.Priority, board.Status)
	output, err := controller.boardCreateUsecase.Create(boardCureateUsecaseInput)
	if err != nil {
		return nil, err
	}
	boardCreatePresenter := presenter_create.NewBoardAllGetPresenter(*output)
	return boardCreatePresenter.Build(), nil
}
