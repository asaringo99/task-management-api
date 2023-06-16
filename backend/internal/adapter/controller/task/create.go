package controller

import (
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	presenter_create "github.com/asaringo99/task_management/internal/adapter/presenter/task/create"
	usecase_create "github.com/asaringo99/task_management/internal/application/usecase/task/create"
	"github.com/labstack/echo/v4"
)

type TaskCreateRequest struct {
	Contents string `json:"contents"`
	Priority int    `json:"priority"`
	Boardid  int    `json:"boardid"`
}

func (controller TaskController) Post(c echo.Context) (*presenter_create.TaskCreatePresenterOutputDto, error) {
	userid := authjwt.RetrieveUserIdFromToken(c)
	task := new(TaskCreateRequest)
	if err := c.Bind(task); err != nil {
		return nil, err
	}
	taskCreateUsecaseInput := usecase_create.NewTaskCreateUsecaseInput(userid, task.Boardid, task.Priority, task.Contents)
	output, err := controller.taskCreateUsecase.Create(taskCreateUsecaseInput)
	if err != nil {
		return nil, err
	}
	taskCreatePresenter := presenter_create.NewTaskCreatePresenter(*output)
	return taskCreatePresenter.Build(), nil
}
