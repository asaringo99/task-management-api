package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/patch"
	"github.com/asaringo99/task_management/internal/application/usecase/board/patch/condition"
)

type BoardPatchRepository struct {
	taskPatchGateway gateway.BoardPatchGatewayInterface
}

func NewBoardPatchRepository(gateway gateway.BoardPatchGatewayInterface) BoardPatchRepositoryInterface {
	return &BoardPatchRepository{
		taskPatchGateway: gateway,
	}
}

func (repository *BoardPatchRepository) Patch(input condition.BoardPatchCondition) error {
	repository.taskPatchGateway.Patch(input)
	return nil
}
