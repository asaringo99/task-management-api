package usecase

type TabCreateUsecaseImpl struct {
	interactor TabCreateInteractor
}

func (c *TabCreateUsecaseImpl) Create(input tabCreateUsecaseInput) (*TabCreateUsecaseOutput, error) {
	return c.interactor.create(input)
}

func NewTabCreateUsecase(interactor *TabCreateInteractor) TabCreateUsecaseInputPort {
	return &TabCreateUsecaseImpl{
		interactor: *interactor,
	}
}
