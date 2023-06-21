package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TabFetchRepositoryInterface interface {
	Find(input condition.TabFetchCondition) ([]TabFetchRepositoryOutput, error)
}

type TabFetchRepositoryInput struct {
	Userid domain.Userid
}

type TabFetchRepositoryOutput struct {
	Tabid    domain.Id
	Userid   domain.Userid
	IsActive bool
	Title    domain.Title
}
