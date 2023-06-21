package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/board/patch/condition"

type BoardPatchUsecaseImpl struct {
	interactor BoardPatchInteractor
}

func (c *BoardPatchUsecaseImpl) Patch(input condition.BoardPatchCondition) error {
	return c.interactor.put(input)
}

func NewBoardPatchUsecase(interactor *BoardPatchInteractor) BoardPatchUsecaseInputPort {
	return &BoardPatchUsecaseImpl{
		interactor: *interactor,
	}
}
