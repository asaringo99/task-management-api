package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"

type TaskPutUsecaseImpl struct {
	interactor TaskPutInteractor
}

func (c *TaskPutUsecaseImpl) Put(input condition.TaskPutCondition) error {
	return c.interactor.put(input)
}

func NewTaskPutUsecase(interactor *TaskPutInteractor) TaskPutUsecaseInputPort {
	return &TaskPutUsecaseImpl{
		interactor: *interactor,
	}
}
