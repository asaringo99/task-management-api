package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/tab/patch"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/patch/condition"
)

type TabPatchInteractor struct {
	repository repository.TabPatchRepositoryInterface
}

func NewTabPatchInteractor(repository repository.TabPatchRepositoryInterface) *TabPatchInteractor {
	return &TabPatchInteractor{
		repository: repository,
	}
}

func (interactor *TabPatchInteractor) put(input condition.TabPatchCondition) error {
	err := interactor.repository.Patch(input)
	return err
}
