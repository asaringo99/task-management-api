package controller

import (
	"fmt"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	presenter_create "github.com/asaringo99/task_management/internal/adapter/presenter/board/create"
	usecase_create "github.com/asaringo99/task_management/internal/application/usecase/board/create"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
	"github.com/labstack/echo/v4"
)

type BoardCreateRequest struct {
	Tabid    int    `json:"tabid"`
	Status   string `json:"status"`
	Priority int    `json:"priority"`
}

func (controller BoardController) Post(c echo.Context) (*presenter_create.BoardCreatePresenterOutputDto, error) {
	userid := authjwt.RetrieveUserIdFromToken(c)
	fmt.Println(userid)
	board := new(BoardCreateRequest)
	if err := c.Bind(board); err != nil {
		return nil, err
	}
	fmt.Println(userid, board.Tabid, board.Status, board.Priority)
	boardCureateUsecaseInput := usecase_create.NewBoardCreateUsecaseInput(
		domain.NewUserid(userid),
		domain.NewId(board.Tabid),
		domain.NewPriority(board.Priority),
		domain.NewStatus(board.Status),
	)
	output, err := controller.boardCreateUsecase.Create(boardCureateUsecaseInput)
	if err != nil {
		return nil, err
	}
	boardCreatePresenter := presenter_create.NewBoardAllGetPresenter(*output)
	return boardCreatePresenter.Build(), nil
}
