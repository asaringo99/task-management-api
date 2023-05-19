package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"

type TaskPutGatewayInterface interface {
	Put(input condition.TaskPutCondition) error
}
