package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
	domain "github.com/asaringo99/task_management/internal/domain/entity"
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
	Status   domain.Status
	Contents domain.Contents
	Priority domain.Priority
}
