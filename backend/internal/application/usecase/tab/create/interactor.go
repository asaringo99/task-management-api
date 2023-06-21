package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/tab/create"
)

type TabCreateInteractor struct {
	repository repository.TabCreateRepositoryInterface
}

func NewTabCreateInteractor(repository repository.TabCreateRepositoryInterface) *TabCreateInteractor {
	return &TabCreateInteractor{
		repository: repository,
	}
}

func (interactor *TabCreateInteractor) create(input tabCreateUsecaseInput) (*TabCreateUsecaseOutput, error) {
	value, err := interactor.repository.Create(convert(input))
	if err != nil {
		return nil, err
	}
	output := TabCreateUsecaseOutput{
		value.Tabid,
		value.Userid,
		value.IsActive,
		value.Title,
	}
	return &output, err
}

func convert(input tabCreateUsecaseInput) repository.TabCreateRepositoryInput {
	return repository.TabCreateRepositoryInput{
		Userid:   input.Userid,
		Tabid:    input.Tabid,
		IsActive: input.IsActive,
		Title:    input.Title,
	}
}
