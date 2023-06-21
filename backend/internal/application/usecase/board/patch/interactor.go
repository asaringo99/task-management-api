package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/board/patch"
	"github.com/asaringo99/task_management/internal/application/usecase/board/patch/condition"
)

type BoardPatchInteractor struct {
	repository repository.BoardPatchRepositoryInterface
}

func NewBoardPatchInteractor(repository repository.BoardPatchRepositoryInterface) *BoardPatchInteractor {
	return &BoardPatchInteractor{
		repository: repository,
	}
}

func (interactor *BoardPatchInteractor) put(input condition.BoardPatchCondition) error {
	err := interactor.repository.Patch(input)
	return err
}
