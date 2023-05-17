package gateway

import model "github.com/asaringo99/task_management/internal/adapter/gateway/task"

type TaskAllGetGatewayMock struct{}

func (c *TaskAllGetGatewayMock) Create(input model.TaskModel) error {
	return nil
}

func NewTaskAllGetGatewayMock() TaskCreateGatewayInterface {
	return &TaskAllGetGatewayMock{}
}
