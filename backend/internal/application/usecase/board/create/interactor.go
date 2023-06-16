package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/board/create"
)

type BoardCreateInteractor struct {
	repository repository.BoardCreateRepositoryInterface
}

func NewBoardCreateInteractor(repository repository.BoardCreateRepositoryInterface) *BoardCreateInteractor {
	return &BoardCreateInteractor{
		repository: repository,
	}
}

func (interactor *BoardCreateInteractor) create(input taskCreateUsecaseInput) (*BoardCreateUsecaseOutput, error) {
	value, err := interactor.repository.Create(convert(input))
	if err != nil {
		return nil, err
	}
	output := BoardCreateUsecaseOutput{
		value.Boardid,
		value.Userid,
		value.Priority,
		value.Status,
	}
	return &output, err
}

func convert(input taskCreateUsecaseInput) repository.BoardCreateRepositoryInput {
	return repository.BoardCreateRepositoryInput{
		Userid:   input.Userid,
		Priority: input.Priority,
		Status:   input.Status,
	}
}
