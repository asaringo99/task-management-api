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

func (interactor *TaskCreateInteractor) create(input taskCreateUsecaseInput) (*TaskCreateUsecaseOutput, error) {
	repositoryOutput, err := interactor.repository.Create(convert(input))
	if err != nil {
		return nil, err
	}
	output := TaskCreateUsecaseOutput{
		repositoryOutput.Taskid,
		repositoryOutput.Userid,
		repositoryOutput.Boardid,
		repositoryOutput.Priority,
		repositoryOutput.Contents,
	}
	return &output, nil
}

func convert(input taskCreateUsecaseInput) repository.TaskCreateRepositoryInput {
	return repository.TaskCreateRepositoryInput{
		UserId:   input.Userid,
		Boardid:  input.Boardid,
		Priority: input.Priority,
		Contents: input.Contents,
	}
}
