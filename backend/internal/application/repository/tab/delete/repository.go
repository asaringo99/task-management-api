package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/delete"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/delete/condition"
)

type TabDeleteRepository struct {
	taskDeleteGateway gateway.TabDeleteGatewayInterface
}

func NewTabDeleteRepository(gateway gateway.TabDeleteGatewayInterface) TabDeleteRepositoryInterface {
	return &TabDeleteRepository{
		taskDeleteGateway: gateway,
	}
}

func (repository *TabDeleteRepository) Delete(input condition.TabDeleteCondition) error {
	repository.taskDeleteGateway.Delete(input)
	return nil
}
