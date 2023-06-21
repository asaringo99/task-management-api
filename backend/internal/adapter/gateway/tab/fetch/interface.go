package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"
)

type TabFetchGatewayInterface interface {
	Find(input condition.TabFetchCondition) ([]model.TabModel, error)
}
