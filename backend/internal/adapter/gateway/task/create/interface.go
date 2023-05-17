package gateway

import model "github.com/asaringo99/task_management/internal/adapter/gateway/task"

type TaskCreateGatewayInterface interface {
	Create(input model.TaskModel) error
}
