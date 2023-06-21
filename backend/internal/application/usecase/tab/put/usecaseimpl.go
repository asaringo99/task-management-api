package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/tab/put/condition"

type TabPutUsecaseImpl struct {
	interactor TabPutInteractor
}

func (c *TabPutUsecaseImpl) Put(input condition.TabPutCondition) error {
	return c.interactor.put(input)
}

func NewTabPutUsecase(interactor *TabPutInteractor) TabPutUsecaseInputPort {
	return &TabPutUsecaseImpl{
		interactor: *interactor,
	}
}
