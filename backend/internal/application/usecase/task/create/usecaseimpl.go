package usecase

type TaskCreateUsecaseImpl struct {
	interactor TaskCreateInteractor
}

func (c *TaskCreateUsecaseImpl) Create(input taskCreateUsecaseInput) error {
	return c.interactor.create(input)
}

func NewTaskCreateUsecase(interactor *TaskCreateInteractor) TaskCreateUsecaseInputPort {
	return &TaskCreateUsecaseImpl{
		interactor: *interactor,
	}
}
