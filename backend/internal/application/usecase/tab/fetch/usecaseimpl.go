package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"

type TabFetchUsecaseImpl struct {
	interactor TabFetchInteractor
}

func (c *TabFetchUsecaseImpl) Find(input condition.TabFetchCondition) ([]TabFetchUsecaseOutput, error) {
	return c.interactor.find(input)
}

func NewFetchGetUsecase(interactor *TabFetchInteractor) TabFetchUsecaseInputPort {
	return &TabFetchUsecaseImpl{
		interactor: *interactor,
	}
}
