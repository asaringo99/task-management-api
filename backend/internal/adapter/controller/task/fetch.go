package controller

import (
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	presenter_fetch "github.com/asaringo99/task_management/internal/adapter/presenter/task/fetch"
	condition_get "github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
	"github.com/labstack/echo/v4"
)

func (controller TaskController) Get(c echo.Context) ([]presenter_fetch.TaskFetchPresenterOutputDto, error) {
	userid := authjwt.RetrieveUserIdFromToken(c)
	input := condition_get.NewTaskFetchCondition(userid)
	output, err := controller.taskFetchUsecase.Find(input)
	if err != nil {
		return nil, err
	}
	taskAllGetPresenter := presenter_fetch.NewTaskAllGetPresenter(output)
	return taskAllGetPresenter.Build(), nil
}
