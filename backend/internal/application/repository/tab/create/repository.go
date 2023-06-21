package repository

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/create"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TabCreateRepository struct {
	taskCreateGateway gateway.TabCreateGatewayInterface
}

func (repository *TabCreateRepository) Create(input TabCreateRepositoryInput) (*TabCreateRepositoryOutput, error) {
	value, err := repository.taskCreateGateway.Create(convert(input))
	if err != nil {
		return nil, err
	}
	output := TabCreateRepositoryOutput{
		domain.NewId(value.Id),
		domain.NewUserid(value.Userid),
		value.IsActive,
		domain.NewTitle(value.Title),
	}
	return &output, nil
}

func convert(input TabCreateRepositoryInput) model.TabModel {
	return model.TabModel{
		Id:       input.Tabid.ToValue(),
		Userid:   input.Userid.ToValue(),
		IsActive: input.IsActive,
		Title:    input.Title.ToValue(),
	}
}

func NewTabCreateRepository(gateway gateway.TabCreateGatewayInterface) TabCreateRepositoryInterface {
	return &TabCreateRepository{
		taskCreateGateway: gateway,
	}
}
