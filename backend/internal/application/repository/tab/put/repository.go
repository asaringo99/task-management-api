package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/put"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/put/condition"
)

type TabPutRepository struct {
	taskPutGateway gateway.TabPutGatewayInterface
}

func NewTabPutRepository(gateway gateway.TabPutGatewayInterface) TabPutRepositoryInterface {
	return &TabPutRepository{
		taskPutGateway: gateway,
	}
}

func (repository *TabPutRepository) Put(input condition.TabPutCondition) error {
	repository.taskPutGateway.Put(input)
	return nil
}
