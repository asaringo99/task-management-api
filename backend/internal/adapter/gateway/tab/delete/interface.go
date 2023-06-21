package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/tab/delete/condition"

type TabDeleteGatewayInterface interface {
	Delete(input condition.TabDeleteCondition) error
}
