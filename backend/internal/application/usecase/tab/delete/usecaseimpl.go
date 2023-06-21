package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/tab/delete/condition"

type TabDeleteUsecaseImpl struct {
	interactor TabDeleteInteractor
}

func (c *TabDeleteUsecaseImpl) Delete(input condition.TabDeleteCondition) error {
	return c.interactor.delete(input)
}

func NewTabDeleteUsecase(interactor *TabDeleteInteractor) TabDeleteUsecaseInputPort {
	return &TabDeleteUsecaseImpl{
		interactor: *interactor,
	}
}
