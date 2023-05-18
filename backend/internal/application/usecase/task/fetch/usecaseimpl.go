package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"

type TaskFetchUsecaseImpl struct {
	interactor TaskFetchInteractor
}

func (c *TaskFetchUsecaseImpl) Find(input condition.TaskFetchCondition) ([]TaskFetchUsecaseOutput, error) {
	return c.interactor.find(input)
}

func NewFetchGetUsecase(interactor *TaskFetchInteractor) TaskFetchUsecaseInputPort {
	return &TaskFetchUsecaseImpl{
		interactor: *interactor,
	}
}
