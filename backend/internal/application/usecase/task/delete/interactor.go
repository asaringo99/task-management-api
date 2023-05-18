package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/task/delete"
	"github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"
)

type TaskDeleteInteractor struct {
	repository repository.TaskDeleteRepositoryInterface
}

func NewTaskDeleteInteractor(repository repository.TaskDeleteRepositoryInterface) *TaskDeleteInteractor {
	return &TaskDeleteInteractor{
		repository: repository,
	}
}

func (interactor *TaskDeleteInteractor) create(input condition.TaskDeleteCondition) error {
	err := interactor.repository.Delete(input)
	return err
}
