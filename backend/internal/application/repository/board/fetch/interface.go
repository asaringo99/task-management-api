package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type BoardFetchRepositoryInterface interface {
	Find(input condition.BoardFetchCondition) ([]BoardFetchRepositoryOutput, error)
}

type BoardFetchRepositoryInput struct {
	Userid domain.Userid
}

type BoardFetchRepositoryOutput struct {
	Boardid  domain.Id
	Userid   domain.Userid
	Priority domain.Priority
	Status   domain.Status
}
