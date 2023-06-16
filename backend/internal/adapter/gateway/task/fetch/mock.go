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
			Boardid:  1,
			Priority: 1,
			Contents: "this is the test",
		},
		{
			Id:       2,
			Userid:   1,
			Boardid:  2,
			Priority: 2,
			Contents: "hello",
		},
		{
			Id:       3,
			Userid:   1,
			Boardid:  1,
			Priority: 3,
			Contents: "hello! world",
		},
	}, nil
}

func NewTaskAllGetGatewayMock() TaskFetchGatewayInterface {
	return &TaskAllGetGatewayMock{}
}
