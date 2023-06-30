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

func NewBoardCreateUsecaseInput(userid domain.Userid, tabid domain.Id, priority domain.Priority, status domain.Status) boardCreateUsecaseInput {
	return boardCreateUsecaseInput{
		Userid:   userid,
		Tabid:    tabid,
		Priority: priority,
		Status:   status,
	}
}
