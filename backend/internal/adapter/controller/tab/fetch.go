package controller

import (
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	presenter "github.com/asaringo99/task_management/internal/adapter/presenter/tab/fetch"
	condition "github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"
	"github.com/labstack/echo/v4"
)

func (controller TabController) Get(c echo.Context) ([]presenter.TabFetchPresenterOutputDto, error) {
	userid := authjwt.RetrieveUserIdFromToken(c)
	input := condition.NewTabFetchCondition(userid)
	output, err := controller.tabFetchUsecase.Find(input)
	if err != nil {
		return nil, err
	}
	tabGetPresenter := presenter.NewTabAllGetPresenter(output)
	return tabGetPresenter.Build(), nil
}
