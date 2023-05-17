package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/task/create"
)

type TaskCreateInteractor struct {
	repository repository.TaskCreateRepositoryInterface
}

func NewTaskCreateInteractor(repository repository.TaskCreateRepositoryInterface) *TaskCreateInteractor {
	return &TaskCreateInteractor{
		repository: repository,
	}
}

func (interactor *TaskCreateInteractor) create(input taskCreateUsecaseInput) error {
	err := interactor.repository.Create(convert(input))
	return err
}

func convert(input taskCreateUsecaseInput) repository.TaskCreateRepositoryInput {
	return repository.TaskCreateRepositoryInput{
		UserId:   input.Userid,
		Status:   input.Status,
		Priority: input.Priority,
		Contents: input.Contents,
	}
}
