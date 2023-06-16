package gateway

import model "github.com/asaringo99/task_management/internal/adapter/gateway/task"

type TaskAllGetGatewayMock struct{}

func (c *TaskAllGetGatewayMock) Create(input model.TaskModel) (model.TaskModel, error) {
	return model.TaskModel{}, nil
}

func NewTaskAllGetGatewayMock() TaskCreateGatewayInterface {
	return &TaskAllGetGatewayMock{}
}
