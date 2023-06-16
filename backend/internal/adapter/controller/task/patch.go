package controller

import (
	"fmt"
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	condition "github.com/asaringo99/task_management/internal/application/usecase/task/patch/condition"
	"github.com/labstack/echo/v4"
)

type TaskPatchRequest struct {
	Boardid int `json:"boardid"`
}

func (controller TaskController) Patch(c echo.Context) error {
	userid := authjwt.RetrieveUserIdFromToken(c)
	taskid, _ := strconv.Atoi(c.Param("id"))
	task := new(TaskCreateRequest)
	if err := c.Bind(task); err != nil {
		return err
	}
	fmt.Println(task, userid)
	input := condition.NewTaskPatchCondition(taskid, userid, condition.PatchData{Column: "boardid", Value: task.Boardid})
	err := controller.taskPatchUsecase.Patch(input)
	if err != nil {
		return err
	}
	return nil
}
