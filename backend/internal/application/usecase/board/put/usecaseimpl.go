package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/board/put/condition"

type BoardPutUsecaseImpl struct {
	interactor BoardPutInteractor
}

func (c *BoardPutUsecaseImpl) Put(input condition.BoardPutCondition) error {
	return c.interactor.put(input)
}

func NewBoardPutUsecase(interactor *BoardPutInteractor) BoardPutUsecaseInputPort {
	return &BoardPutUsecaseImpl{
		interactor: *interactor,
	}
}
