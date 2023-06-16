package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/board/delete/condition"

type BoardDeleteGatewayInterface interface {
	Delete(input condition.BoardDeleteCondition) error
}
