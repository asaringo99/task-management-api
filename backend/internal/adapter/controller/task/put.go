package controller

import (
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	condition_put "github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"
	"github.com/labstack/echo/v4"
)

type TaskPutRequest struct {
	Contents string `json:"contents"`
	Priority int    `json:"priority"`
	Boardid  int    `json:"boardid"`
}

func (controller TaskController) Put(c echo.Context) error {
	userid := authjwt.RetrieveUserIdFromToken(c)
	taskid, _ := strconv.Atoi(c.Param("id"))
	task := new(TaskCreateRequest)
	if err := c.Bind(task); err != nil {
		return err
	}
	input := condition_put.NewTaskPutCondition(taskid, userid, task.Boardid, task.Priority, task.Contents)
	err := controller.taskPutUsecase.Put(input)
	if err != nil {
		return err
	}
	return nil
}
