package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/delete"
	"github.com/asaringo99/task_management/internal/application/usecase/board/delete/condition"
)

type BoardDeleteRepository struct {
	taskDeleteGateway gateway.BoardDeleteGatewayInterface
}

func NewBoardDeleteRepository(gateway gateway.BoardDeleteGatewayInterface) BoardDeleteRepositoryInterface {
	return &BoardDeleteRepository{
		taskDeleteGateway: gateway,
	}
}

func (repository *BoardDeleteRepository) Delete(input condition.BoardDeleteCondition) error {
	repository.taskDeleteGateway.Delete(input)
	return nil
}
