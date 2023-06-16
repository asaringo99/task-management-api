package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/put"
	"github.com/asaringo99/task_management/internal/application/usecase/board/put/condition"
)

type BoardPutRepository struct {
	taskPutGateway gateway.BoardPutGatewayInterface
}

func NewBoardPutRepository(gateway gateway.BoardPutGatewayInterface) BoardPutRepositoryInterface {
	return &BoardPutRepository{
		taskPutGateway: gateway,
	}
}

func (repository *BoardPutRepository) Put(input condition.BoardPutCondition) error {
	repository.taskPutGateway.Put(input)
	return nil
}
