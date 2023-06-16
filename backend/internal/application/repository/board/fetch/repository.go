package repository

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/fetch"
	"github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type BoardFetchRepository struct {
	taskFetchGateway gateway.BoardFetchGatewayInterface
}

func NewBoardFetchRepository(gateway gateway.BoardFetchGatewayInterface) BoardFetchRepositoryInterface {
	return &BoardFetchRepository{
		taskFetchGateway: gateway,
	}
}

func (repository *BoardFetchRepository) Find(input condition.BoardFetchCondition) ([]BoardFetchRepositoryOutput, error) {
	response, err := repository.taskFetchGateway.Find(input)
	if err != nil {
		return nil, err
	}
	ret := []BoardFetchRepositoryOutput{}
	for _, r := range response {
		ret = append(ret, convert(r))
	}
	return ret, nil
}

func convert(value model.BoardModel) BoardFetchRepositoryOutput {
	return BoardFetchRepositoryOutput{
		domain.NewId(value.Id),
		domain.NewUserid(value.Userid),
		domain.NewPriority(value.Priority),
		domain.NewStatus(value.Status),
	}
}
