package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/task/patch"
	"github.com/asaringo99/task_management/internal/application/usecase/task/patch/condition"
)

type TaskPatchInteractor struct {
	repository repository.TaskPatchRepositoryInterface
}

func NewTaskPatchInteractor(repository repository.TaskPatchRepositoryInterface) *TaskPatchInteractor {
	return &TaskPatchInteractor{
		repository: repository,
	}
}

func (interactor *TaskPatchInteractor) put(input condition.TaskPatchCondition) error {
	err := interactor.repository.Patch(input)
	return err
}
