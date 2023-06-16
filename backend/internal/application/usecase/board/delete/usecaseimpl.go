package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/board/delete/condition"

type BoardDeleteUsecaseImpl struct {
	interactor BoardDeleteInteractor
}

func (c *BoardDeleteUsecaseImpl) Delete(input condition.BoardDeleteCondition) error {
	return c.interactor.create(input)
}

func NewBoardDeleteUsecase(interactor *BoardDeleteInteractor) BoardDeleteUsecaseInputPort {
	return &BoardDeleteUsecaseImpl{
		interactor: *interactor,
	}
}
