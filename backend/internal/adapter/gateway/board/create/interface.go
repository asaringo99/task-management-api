package gateway

import model "github.com/asaringo99/task_management/internal/adapter/gateway/board"

type BoardCreateGatewayInterface interface {
	Create(input model.BoardModel) (model.BoardModel, error)
}
