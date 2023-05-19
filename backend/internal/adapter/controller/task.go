package controller

import (
	"strconv"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	presenter "github.com/asaringo99/task_management/internal/adapter/presenter/task"
	usecase_create "github.com/asaringo99/task_management/internal/application/usecase/task/create"
	usecase_delete "github.com/asaringo99/task_management/internal/application/usecase/task/delete"
	condition_delete "github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"
	usecase_get "github.com/asaringo99/task_management/internal/application/usecase/task/fetch"
	condition_get "github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
	usecase_put "github.com/asaringo99/task_management/internal/application/usecase/task/put"
	condition_put "github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"
	"github.com/labstack/echo/v4"
)

type TaskController struct {
	taskFetchUsecase  usecase_get.TaskFetchUsecaseInputPort
	taskCreateUsecase usecase_create.TaskCreateUsecaseInputPort
	taskDeleteUsecase usecase_delete.TaskDeleteUsecaseInputPort
	taskPutUsecase    usecase_put.TaskPutUsecaseInputPort
}

func (controller TaskController) Get(c echo.Context) ([]presenter.TaskFetchPresenterOutputDto, error) {
	userid := authjwt.RetrieveUserIdFromToken(c)
	input := condition_get.NewTaskFetchCondition(userid)
	output, err := controller.taskFetchUsecase.Find(input)
	if err != nil {
		return nil, err
	}
	taskAllGetPresenter := presenter.NewTaskAllGetPresenter(output)
	return taskAllGetPresenter.Build(), nil
}

func (controller TaskController) Post(c echo.Context) error {
	// TODO: Bindで受け取る
	userid := authjwt.RetrieveUserIdFromToken(c)
	status := c.FormValue("status")
	contents := c.FormValue("contents")
	priority, _ := strconv.Atoi(c.FormValue("priority"))
	task := usecase_create.NewTaskCreateUsecaseInput(userid, status, priority, contents)
	err := controller.taskCreateUsecase.Create(task)
	if err != nil {
		return err
	}
	return nil
}

func (controller TaskController) Put(c echo.Context) error {
	// TODO: Bindで受け取る
	userid := authjwt.RetrieveUserIdFromToken(c)
	taskid, _ := strconv.Atoi(c.Param("id"))
	status := c.FormValue("status")
	contents := c.FormValue("contents")
	priority, _ := strconv.Atoi(c.FormValue("priority"))
	input := condition_put.NewTaskPutCondition(taskid, userid, status, priority, contents)
	err := controller.taskPutUsecase.Put(input)
	if err != nil {
		return err
	}
	return nil
}

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

func NewManagementController(
	taskAllGetUsecase usecase_get.TaskFetchUsecaseInputPort,
	taskCreateUsecase usecase_create.TaskCreateUsecaseInputPort,
	taskDeleteUsecase usecase_delete.TaskDeleteUsecaseInputPort,
	taskPutUsecase usecase_put.TaskPutUsecaseInputPort,
) *TaskController {
	return &TaskController{
		taskFetchUsecase:  taskAllGetUsecase,
		taskCreateUsecase: taskCreateUsecase,
		taskDeleteUsecase: taskDeleteUsecase,
		taskPutUsecase:    taskPutUsecase,
	}
}
