package repository

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/create"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type BoardCreateRepository struct {
	taskCreateGateway gateway.BoardCreateGatewayInterface
}

func (repository *BoardCreateRepository) Create(input BoardCreateRepositoryInput) (*BoardCreateRepositoryOutput, error) {
	value, err := repository.taskCreateGateway.Create(convert(input))
	if err != nil {
		return nil, err
	}
	output := BoardCreateRepositoryOutput{
		domain.NewId(value.Id),
		domain.NewUserid(value.Userid),
		domain.NewId(value.Tabid),
		domain.NewPriority(value.Priority),
		domain.NewStatus(value.Status),
	}
	return &output, nil
}

func convert(input BoardCreateRepositoryInput) model.BoardModel {
	return model.BoardModel{
		Userid:   input.Userid.ToValue(),
		Tabid:    input.Tabid.ToValue(),
		Priority: input.Priority.ToValue(),
		Status:   input.Status.ToValue(),
	}
}

func NewBoardCreateRepository(gateway gateway.BoardCreateGatewayInterface) BoardCreateRepositoryInterface {
	return &BoardCreateRepository{
		taskCreateGateway: gateway,
	}
}
