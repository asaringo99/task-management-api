package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/delete"
	"github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"
)

type TaskDeleteRepository struct {
	taskDeleteGateway gateway.TaskDeleteGatewayInterface
}

func NewTaskDeleteRepository(gateway gateway.TaskDeleteGatewayInterface) TaskDeleteRepositoryInterface {
	return &TaskDeleteRepository{
		taskDeleteGateway: gateway,
	}
}

func (repository *TaskDeleteRepository) Delete(input condition.TaskDeleteCondition) error {
	repository.taskDeleteGateway.Delete(input)
	return nil
}
