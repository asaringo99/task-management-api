package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/tab/delete"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/delete/condition"
)

type TabDeleteInteractor struct {
	repository repository.TabDeleteRepositoryInterface
}

func NewTabDeleteInteractor(repository repository.TabDeleteRepositoryInterface) *TabDeleteInteractor {
	return &TabDeleteInteractor{
		repository: repository,
	}
}

func (interactor *TabDeleteInteractor) delete(input condition.TabDeleteCondition) error {
	err := interactor.repository.Delete(input)
	return err
}
