package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/put"
	"github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"
)

type TaskPutRepository struct {
	taskPutGateway gateway.TaskPutGatewayInterface
}

func NewTaskPutRepository(gateway gateway.TaskPutGatewayInterface) TaskPutRepositoryInterface {
	return &TaskPutRepository{
		taskPutGateway: gateway,
	}
}

func (repository *TaskPutRepository) Put(input condition.TaskPutCondition) error {
	repository.taskPutGateway.Put(input)
	return nil
}
