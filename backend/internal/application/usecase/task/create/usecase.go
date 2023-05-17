package usecase

import (
	domain "github.com/asaringo99/task_management/internal/domain/entity"
)

type TaskCreateUsecaseInputPort interface {
	Create(taskCreateUsecaseInput) error
}

type taskCreateUsecaseInput struct {
	Userid   domain.Userid
	Status   domain.Status
	Priority domain.Priority
	Contents domain.Contents
}

func NewTaskCreateUsecaseInput(userid int, status string, priority int, contents string) taskCreateUsecaseInput {
	return taskCreateUsecaseInput{
		Userid:   domain.NewUserid(userid),
		Status:   domain.NewStatus(status),
		Priority: domain.NewPriority(priority),
		Contents: domain.NewContents(contents),
	}
}
