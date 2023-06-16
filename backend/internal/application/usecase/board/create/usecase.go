package usecase

import (
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type BoardCreateUsecaseInputPort interface {
	Create(taskCreateUsecaseInput) (*BoardCreateUsecaseOutput, error)
}

type taskCreateUsecaseInput struct {
	Userid   domain.Userid
	Priority domain.Priority
	Status   domain.Status
}

type BoardCreateUsecaseOutput struct {
	Boardid  domain.Id
	Userid   domain.Userid
	Priority domain.Priority
	Status   domain.Status
}

func NewBoardCreateUsecaseInput(userid int, priority int, status string) taskCreateUsecaseInput {
	return taskCreateUsecaseInput{
		Userid:   domain.NewUserid(userid),
		Priority: domain.NewPriority(priority),
		Status:   domain.NewStatus(status),
	}
}
