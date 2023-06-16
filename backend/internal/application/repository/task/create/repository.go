package repository

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/create"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TaskCreateRepository struct {
	taskCreateGateway gateway.TaskCreateGatewayInterface
}

func NewTaskCreateRepository(gateway gateway.TaskCreateGatewayInterface) TaskCreateRepositoryInterface {
	return &TaskCreateRepository{
		taskCreateGateway: gateway,
	}
}

func (repository *TaskCreateRepository) Create(input TaskCreateRepositoryInput) (*TaskCreateRepositoryOutput, error) {
	value, err := repository.taskCreateGateway.Create(convert(input))
	if err != nil {
		return nil, err
	}
	output := TaskCreateRepositoryOutput{
		domain.NewTaskid(value.Id),
		domain.NewUserid(value.Userid),
		domain.NewId(value.Boardid),
		domain.NewContents(value.Contents),
		domain.NewPriority(value.Priority),
	}
	return &output, nil
}

func convert(input TaskCreateRepositoryInput) model.TaskModel {
	return model.TaskModel{
		Userid:   input.UserId.ToValue(),
		Boardid:  input.Boardid.ToValue(),
		Priority: input.Priority.ToValue(),
		Contents: input.Contents.ToValue(),
	}
}
