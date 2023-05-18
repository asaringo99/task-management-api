package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"

type TaskDeleteGatewayInterface interface {
	Delete(input condition.TaskDeleteCondition) error
}
