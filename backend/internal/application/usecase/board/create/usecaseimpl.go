package usecase

type BoardCreateUsecaseImpl struct {
	interactor BoardCreateInteractor
}

func (c *BoardCreateUsecaseImpl) Create(input boardCreateUsecaseInput) (*BoardCreateUsecaseOutput, error) {
	return c.interactor.create(input)
}

func NewBoardCreateUsecase(interactor *BoardCreateInteractor) BoardCreateUsecaseInputPort {
	return &BoardCreateUsecaseImpl{
		interactor: *interactor,
	}
}
