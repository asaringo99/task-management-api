package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/board/fetch"
	"github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"
)

type BoardFetchInteractor struct {
	repository repository.BoardFetchRepositoryInterface
}

func NewBoardFetchInteractor(repository repository.BoardFetchRepositoryInterface) *BoardFetchInteractor {
	return &BoardFetchInteractor{
		repository: repository,
	}
}

func (interactor *BoardFetchInteractor) find(input condition.BoardFetchCondition) ([]BoardFetchUsecaseOutput, error) {
	response, err := interactor.repository.Find(input)
	if err != nil {
		return nil, err
	}
	ret := []BoardFetchUsecaseOutput{}
	for _, res := range response {
		ret = append(ret, convert(res))
	}
	return ret, nil
}

func convert(value repository.BoardFetchRepositoryOutput) BoardFetchUsecaseOutput {
	return BoardFetchUsecaseOutput{
		Boardid:  value.Boardid,
		Userid:   value.Userid,
		Tabid:    value.Tabid,
		Priority: value.Priority,
		Status:   value.Status,
	}
}
