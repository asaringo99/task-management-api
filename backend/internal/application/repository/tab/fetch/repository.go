package repository

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/fetch"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TabFetchRepository struct {
	taskFetchGateway gateway.TabFetchGatewayInterface
}

func NewTabFetchRepository(gateway gateway.TabFetchGatewayInterface) TabFetchRepositoryInterface {
	return &TabFetchRepository{
		taskFetchGateway: gateway,
	}
}

func (repository *TabFetchRepository) Find(input condition.TabFetchCondition) ([]TabFetchRepositoryOutput, error) {
	response, err := repository.taskFetchGateway.Find(input)
	if err != nil {
		return nil, err
	}
	ret := []TabFetchRepositoryOutput{}
	for _, r := range response {
		ret = append(ret, convert(r))
	}
	return ret, nil
}

func convert(value model.TabModel) TabFetchRepositoryOutput {
	return TabFetchRepositoryOutput{
		domain.NewId(value.Id),
		domain.NewUserid(value.Userid),
		value.IsActive,
		domain.NewTitle(value.Title),
	}
}
