package usecase

import (
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type BoardCreateUsecaseInputPort interface {
	Create(boardCreateUsecaseInput) (*BoardCreateUsecaseOutput, error)
}

type boardCreateUsecaseInput struct {
	Userid   domain.Userid
	Tabid    domain.Id
	Priority domain.Priority
	Status   domain.Status
}

type BoardCreateUsecaseOutput struct {
	Boardid  domain.Id
	Userid   domain.Userid
	Tabid    domain.Id
	Priority domain.Priority
	Status   domain.Status
}

func NewBoardCreateUsecaseInput(userid int, tabid int, priority int, status string) boardCreateUsecaseInput {
	return boardCreateUsecaseInput{
		Userid:   domain.NewUserid(userid),
		Tabid:    domain.NewId(tabid),
		Priority: domain.NewPriority(priority),
		Status:   domain.NewStatus(status),
	}
}
