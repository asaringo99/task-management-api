package controller

import (
	usecase_create "github.com/asaringo99/task_management/internal/application/usecase/task/create"
	usecase_delete "github.com/asaringo99/task_management/internal/application/usecase/task/delete"
	usecase_get "github.com/asaringo99/task_management/internal/application/usecase/task/fetch"
	usecase_patch "github.com/asaringo99/task_management/internal/application/usecase/task/patch"
	usecase_put "github.com/asaringo99/task_management/internal/application/usecase/task/put"
)

type TaskController struct {
	taskFetchUsecase  usecase_get.TaskFetchUsecaseInputPort
	taskCreateUsecase usecase_create.TaskCreateUsecaseInputPort
	taskDeleteUsecase usecase_delete.TaskDeleteUsecaseInputPort
	taskPutUsecase    usecase_put.TaskPutUsecaseInputPort
	taskPatchUsecase  usecase_patch.TaskPatchUsecaseInputPort
}

func NewManagementController(
	taskAllGetUsecase usecase_get.TaskFetchUsecaseInputPort,
	taskCreateUsecase usecase_create.TaskCreateUsecaseInputPort,
	taskDeleteUsecase usecase_delete.TaskDeleteUsecaseInputPort,
	taskPutUsecase usecase_put.TaskPutUsecaseInputPort,
	taskPatchUsecase usecase_patch.TaskPatchUsecaseInputPort,
) *TaskController {
	return &TaskController{
		taskFetchUsecase:  taskAllGetUsecase,
		taskCreateUsecase: taskCreateUsecase,
		taskDeleteUsecase: taskDeleteUsecase,
		taskPutUsecase:    taskPutUsecase,
		taskPatchUsecase:  taskPatchUsecase,
	}
}
