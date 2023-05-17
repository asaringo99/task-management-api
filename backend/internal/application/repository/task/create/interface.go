package repository

import (
	domain "github.com/asaringo99/task_management/internal/domain/entity"
)

type TaskCreateRepositoryInterface interface {
	Create(TaskCreateRepositoryInput) error
}

type TaskCreateRepositoryInput struct {
	UserId   domain.Userid
	UserName domain.Username
	Status   domain.Status
	Contents domain.Contents
	Priority domain.Priority
}
