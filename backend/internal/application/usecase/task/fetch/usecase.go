package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/entity"
)

type TaskFetchUsecaseInputPort interface {
	Find(input condition.TaskFetchCondition) ([]TaskFetchUsecaseOutput, error)
}

type TaskFetchUsecaseOutput struct {
	Taskid   domain.Taskid
	Userid   domain.Userid
	Status   domain.Status
	Priority domain.Priority
	Contents domain.Contents
}
