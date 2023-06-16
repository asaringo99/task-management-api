package controller

import (
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	condition_delete "github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"
	"github.com/labstack/echo/v4"
)

func (controller TaskController) Delete(c echo.Context) error {
	userid := authjwt.RetrieveUserIdFromToken(c)
	taskid, _ := strconv.Atoi(c.Param("id"))
	task := condition_delete.NewTaskDeleteCondition(userid, taskid)
	err := controller.taskDeleteUsecase.Delete(task)
	if err != nil {
		return err
	}
	return nil
}
