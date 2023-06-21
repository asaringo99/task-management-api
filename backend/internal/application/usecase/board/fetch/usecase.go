package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type BoardFetchUsecaseInputPort interface {
	Find(input condition.BoardFetchCondition) ([]BoardFetchUsecaseOutput, error)
}

type BoardFetchUsecaseOutput struct {
	Boardid  domain.Id
	Userid   domain.Userid
	Tabid    domain.Id
	Priority domain.Priority
	Status   domain.Status
}
