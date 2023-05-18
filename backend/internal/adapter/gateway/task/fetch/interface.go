package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
)

type TaskFetchGatewayInterface interface {
	Find(input condition.TaskFetchCondition) ([]model.TaskModel, error)
}
