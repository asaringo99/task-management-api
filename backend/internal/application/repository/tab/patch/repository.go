package repository

import (
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/patch"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/patch/condition"
)

type TabPatchRepository struct {
	taskPatchGateway gateway.TabPatchGatewayInterface
}

func NewTabPatchRepository(gateway gateway.TabPatchGatewayInterface) TabPatchRepositoryInterface {
	return &TabPatchRepository{
		taskPatchGateway: gateway,
	}
}

func (repository *TabPatchRepository) Patch(input condition.TabPatchCondition) error {
	repository.taskPatchGateway.Patch(input)
	return nil
}
