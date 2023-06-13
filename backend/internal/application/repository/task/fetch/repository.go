package repository

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/fetch"
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TaskFetchRepository struct {
	taskFetchGateway gateway.TaskFetchGatewayInterface
}

func NewTaskFetchRepository(gateway gateway.TaskFetchGatewayInterface) TaskFetchRepositoryInterface {
	return &TaskFetchRepository{
		taskFetchGateway: gateway,
	}
}

func (repository *TaskFetchRepository) Find(input condition.TaskFetchCondition) ([]TaskFetchRepositoryOutput, error) {
	response, err := repository.taskFetchGateway.Find(input)
	if err != nil {
		return nil, err
	}
	ret := []TaskFetchRepositoryOutput{}
	for _, r := range response {
		ret = append(ret, convert(r))
	}
	return ret, nil
}

func convert(value model.TaskModel) TaskFetchRepositoryOutput {
	return TaskFetchRepositoryOutput{
		domain.NewTaskid(value.Id),
		domain.NewUserid(value.Userid),
		domain.NewStatus(value.Status),
		domain.NewContents(value.Contents),
		domain.NewPriority(value.Priority),
	}
}
