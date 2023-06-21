package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TabFetchUsecaseInputPort interface {
	Find(input condition.TabFetchCondition) ([]TabFetchUsecaseOutput, error)
}

type TabFetchUsecaseOutput struct {
	Tabid    domain.Id
	Userid   domain.Userid
	IsActive bool
	Title    domain.Title
}
