package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	"github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"
)

type BoardFetchGatewayInterface interface {
	Find(input condition.BoardFetchCondition) ([]model.BoardModel, error)
}
