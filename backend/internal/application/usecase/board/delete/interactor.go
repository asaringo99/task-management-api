package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/board/delete"
	"github.com/asaringo99/task_management/internal/application/usecase/board/delete/condition"
)

type BoardDeleteInteractor struct {
	repository repository.BoardDeleteRepositoryInterface
}

func NewBoardDeleteInteractor(repository repository.BoardDeleteRepositoryInterface) *BoardDeleteInteractor {
	return &BoardDeleteInteractor{
		repository: repository,
	}
}

func (interactor *BoardDeleteInteractor) delete(input condition.BoardDeleteCondition) error {
	err := interactor.repository.Delete(input)
	return err
}
