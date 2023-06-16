package controller

import (
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	presenter "github.com/asaringo99/task_management/internal/adapter/presenter/board/fetch"
	condition "github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"
	"github.com/labstack/echo/v4"
)

func (controller BoardController) Get(c echo.Context) ([]presenter.BoardFetchPresenterOutputDto, error) {
	userid := authjwt.RetrieveUserIdFromToken(c)
	input := condition.NewBoardFetchCondition(userid)
	output, err := controller.boardFetchUsecase.Find(input)
	if err != nil {
		return nil, err
	}
	boardGetPresenter := presenter.NewBoardAllGetPresenter(output)
	return boardGetPresenter.Build(), nil
}
