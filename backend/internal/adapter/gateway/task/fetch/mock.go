package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
)

type TaskAllGetGatewayMock struct{}

func (c *TaskAllGetGatewayMock) Find(input condition.TaskFetchCondition) ([]model.TaskModel, error) {
	return []model.TaskModel{
		{
			Id:       1,
			Userid:   1,
			Status:   "Ready",
			Priority: 1,
			Contents: "this is the test",
		},
		{
			Id:       2,
			Userid:   1,
			Status:   "Todo",
			Priority: 2,
			Contents: "hello",
		},
		{
			Id:       3,
			Userid:   1,
			Status:   "Closed",
			Priority: 3,
			Contents: "hello! world",
		},
	}, nil
}

func NewTaskAllGetGatewayMock() TaskFetchGatewayInterface {
	return &TaskAllGetGatewayMock{}
}
