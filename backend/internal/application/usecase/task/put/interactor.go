package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/task/put"
	"github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"
)

type TaskPutInteractor struct {
	repository repository.TaskPutRepositoryInterface
}

func NewTaskPutInteractor(repository repository.TaskPutRepositoryInterface) *TaskPutInteractor {
	return &TaskPutInteractor{
		repository: repository,
	}
}

func (interactor *TaskPutInteractor) put(input condition.TaskPutCondition) error {
	err := interactor.repository.Put(input)
	return err
}
