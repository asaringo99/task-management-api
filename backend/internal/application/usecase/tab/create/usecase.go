package usecase

import (
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TabCreateUsecaseInputPort interface {
	Create(tabCreateUsecaseInput) (*TabCreateUsecaseOutput, error)
}

type tabCreateUsecaseInput struct {
	Tabid    domain.Id
	Userid   domain.Userid
	IsActive bool
	Title    domain.Title
}

type TabCreateUsecaseOutput struct {
	Tabid    domain.Id
	Userid   domain.Userid
	IsActive bool
	Title    domain.Title
}

func NewTabCreateUsecaseInput(userid int, tabid int, isActive bool, title string) tabCreateUsecaseInput {
	return tabCreateUsecaseInput{
		Userid:   domain.NewUserid(userid),
		Tabid:    domain.NewId(tabid),
		IsActive: isActive,
		Title:    domain.NewTitle(title),
	}
}
