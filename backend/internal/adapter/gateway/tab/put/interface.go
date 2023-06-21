package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/tab/put/condition"

type TabPutGatewayInterface interface {
	Put(input condition.TabPutCondition) error
}
