package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/tab/fetch"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"
)

type TabFetchInteractor struct {
	repository repository.TabFetchRepositoryInterface
}

func NewTabFetchInteractor(repository repository.TabFetchRepositoryInterface) *TabFetchInteractor {
	return &TabFetchInteractor{
		repository: repository,
	}
}

func (interactor *TabFetchInteractor) find(input condition.TabFetchCondition) ([]TabFetchUsecaseOutput, error) {
	response, err := interactor.repository.Find(input)
	if err != nil {
		return nil, err
	}
	ret := []TabFetchUsecaseOutput{}
	for _, res := range response {
		ret = append(ret, convert(res))
	}
	return ret, nil
}

func convert(value repository.TabFetchRepositoryOutput) TabFetchUsecaseOutput {
	return TabFetchUsecaseOutput{
		Tabid:    value.Tabid,
		Userid:   value.Userid,
		IsActive: value.IsActive,
		Title:    value.Title,
	}
}
