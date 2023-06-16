package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/patch"
	"github.com/asaringo99/task_management/internal/application/usecase/task/patch/condition"
)

type TaskPatchRepository struct {
	taskPatchGateway gateway.TaskPatchGatewayInterface
}

func NewTaskPatchRepository(gateway gateway.TaskPatchGatewayInterface) TaskPatchRepositoryInterface {
	return &TaskPatchRepository{
		taskPatchGateway: gateway,
	}
}

func (repository *TaskPatchRepository) Patch(input condition.TaskPatchCondition) error {
	repository.taskPatchGateway.Patch(input)
	return nil
}
