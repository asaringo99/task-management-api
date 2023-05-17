package repository

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/create"
)

type TaskCreateRepository struct {
	taskCreateGateway gateway.TaskCreateGatewayInterface
}

func NewTaskCreateRepository(gateway gateway.TaskCreateGatewayInterface) TaskCreateRepositoryInterface {
	return &TaskCreateRepository{
		taskCreateGateway: gateway,
	}
}

func (repository *TaskCreateRepository) Create(input TaskCreateRepositoryInput) error {
	repository.taskCreateGateway.Create(convert(input))
	return nil
}

func convert(input TaskCreateRepositoryInput) model.TaskModel {
	return model.TaskModel{
		Userid:   input.UserId.ToValue(),
		Status:   input.Status.ToValue(),
		Priority: input.Priority.ToValue(),
		Contents: input.Contents.ToValue(),
	}
}
