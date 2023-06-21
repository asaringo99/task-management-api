package usecase

import "github.com/asaringo99/task_management/internal/application/usecase/tab/patch/condition"

type TabPatchUsecaseImpl struct {
	interactor TabPatchInteractor
}

func (c *TabPatchUsecaseImpl) Patch(input condition.TabPatchCondition) error {
	return c.interactor.put(input)
}

func NewTabPatchUsecase(interactor *TabPatchInteractor) TabPatchUsecaseInputPort {
	return &TabPatchUsecaseImpl{
		interactor: *interactor,
	}
}
