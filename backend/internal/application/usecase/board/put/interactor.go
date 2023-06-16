package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/board/put"
	"github.com/asaringo99/task_management/internal/application/usecase/board/put/condition"
)

type BoardPutInteractor struct {
	repository repository.BoardPutRepositoryInterface
}

func NewBoardPutInteractor(repository repository.BoardPutRepositoryInterface) *BoardPutInteractor {
	return &BoardPutInteractor{
		repository: repository,
	}
}

func (interactor *BoardPutInteractor) put(input condition.BoardPutCondition) error {
	err := interactor.repository.Put(input)
	return err
}
