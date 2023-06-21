package gateway

import model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"

type TabCreateGatewayInterface interface {
	Create(input model.TabModel) (model.TabModel, error)
}
