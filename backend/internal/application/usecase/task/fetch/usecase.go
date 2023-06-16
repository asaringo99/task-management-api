package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TaskFetchUsecaseInputPort interface {
	Find(input condition.TaskFetchCondition) ([]TaskFetchUsecaseOutput, error)
}

type TaskFetchUsecaseOutput struct {
	Taskid   domain.Taskid
	Userid   domain.Userid
	Boardid  domain.Id
	Priority domain.Priority
	Contents domain.Contents
}
