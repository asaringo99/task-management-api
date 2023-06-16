package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/task/patch/condition"

type TaskPatchGatewayInterface interface {
	Patch(input condition.TaskPatchCondition) error
}
