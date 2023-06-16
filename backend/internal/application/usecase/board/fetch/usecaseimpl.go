package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"

type BoardFetchUsecaseImpl struct {
	interactor BoardFetchInteractor
}

func (c *BoardFetchUsecaseImpl) Find(input condition.BoardFetchCondition) ([]BoardFetchUsecaseOutput, error) {
	return c.interactor.find(input)
}

func NewFetchGetUsecase(interactor *BoardFetchInteractor) BoardFetchUsecaseInputPort {
	return &BoardFetchUsecaseImpl{
		interactor: *interactor,
	}
}
