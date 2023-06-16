package repository

import (
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TaskCreateRepositoryInterface interface {
	Create(TaskCreateRepositoryInput) (*TaskCreateRepositoryOutput, error)
}

type TaskCreateRepositoryInput struct {
	UserId   domain.Userid
	Boardid  domain.Id
	Contents domain.Contents
	Priority domain.Priority
}

type TaskCreateRepositoryOutput struct {
	Taskid   domain.Taskid
	Userid   domain.Userid
	Boardid  domain.Id
	Contents domain.Contents
	Priority domain.Priority
}
