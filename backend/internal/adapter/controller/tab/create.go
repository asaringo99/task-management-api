package controller

import (
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	presenter_create "github.com/asaringo99/task_management/internal/adapter/presenter/tab/create"
	usecase_create "github.com/asaringo99/task_management/internal/application/usecase/tab/create"
	"github.com/labstack/echo/v4"
)

type TabCreateRequest struct {
	Tabid    int    `json:"tabid"`
	IsActive bool   `json:"isactive"`
	Title    string `json:"title"`
}

func (controller TabController) Post(c echo.Context) (*presenter_create.TabCreatePresenterOutputDto, error) {
	userid := authjwt.RetrieveUserIdFromToken(c)
	tab := new(TabCreateRequest)
	if err := c.Bind(tab); err != nil {
		return nil, err
	}
	tabCureateUsecaseInput := usecase_create.NewTabCreateUsecaseInput(userid, tab.Tabid, tab.IsActive, tab.Title)
	output, err := controller.tabCreateUsecase.Create(tabCureateUsecaseInput)
	if err != nil {
		return nil, err
	}
	tabCreatePresenter := presenter_create.NewTabAllGetPresenter(*output)
	return tabCreatePresenter.Build(), nil
}
