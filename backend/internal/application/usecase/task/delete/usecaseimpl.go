package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"

type TaskDeleteUsecaseImpl struct {
	interactor TaskDeleteInteractor
}

func (c *TaskDeleteUsecaseImpl) Delete(input condition.TaskDeleteCondition) error {
	return c.interactor.create(input)
}

func NewTaskDeleteUsecase(interactor *TaskDeleteInteractor) TaskDeleteUsecaseInputPort {
	return &TaskDeleteUsecaseImpl{
		interactor: *interactor,
	}
}
