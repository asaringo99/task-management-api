package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TaskFetchRepositoryInterface interface {
	Find(input condition.TaskFetchCondition) ([]TaskFetchRepositoryOutput, error)
}

type TaskFetchRepositoryInput struct {
	Userid domain.Userid
}

type TaskFetchRepositoryOutput struct {
	Taskid   domain.Taskid
	Userid   domain.Userid
	Boardid  domain.Id
	Contents domain.Contents
	Priority domain.Priority
}
