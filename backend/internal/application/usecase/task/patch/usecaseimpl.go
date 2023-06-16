package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/task/patch/condition"

type TaskPatchUsecaseImpl struct {
	interactor TaskPatchInteractor
}

func (c *TaskPatchUsecaseImpl) Patch(input condition.TaskPatchCondition) error {
	return c.interactor.put(input)
}

func NewTaskPatchUsecase(interactor *TaskPatchInteractor) TaskPatchUsecaseInputPort {
	return &TaskPatchUsecaseImpl{
		interactor: *interactor,
	}
}
