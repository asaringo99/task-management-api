package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/tab/put"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/put/condition"
)

type TabPutInteractor struct {
	repository repository.TabPutRepositoryInterface
}

func NewTabPutInteractor(repository repository.TabPutRepositoryInterface) *TabPutInteractor {
	return &TabPutInteractor{
		repository: repository,
	}
}

func (interactor *TabPutInteractor) put(input condition.TabPutCondition) error {
	err := interactor.repository.Put(input)
	return err
}
