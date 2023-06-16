package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/board/put/condition"

type BoardPutGatewayInterface interface {
	Put(input condition.BoardPutCondition) error
}
