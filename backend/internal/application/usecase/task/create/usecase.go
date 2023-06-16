package usecase

import (
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TaskCreateUsecaseInputPort interface {
	Create(taskCreateUsecaseInput) (*TaskCreateUsecaseOutput, error)
}

type taskCreateUsecaseInput struct {
	Userid   domain.Userid
	Boardid  domain.Id
	Priority domain.Priority
	Contents domain.Contents
}

type TaskCreateUsecaseOutput struct {
	Taskid   domain.Taskid
	Userid   domain.Userid
	Boardid  domain.Id
	Priority domain.Priority
	Contents domain.Contents
}

func NewTaskCreateUsecaseInput(userid int, boardid int, priority int, contents string) taskCreateUsecaseInput {
	return taskCreateUsecaseInput{
		Userid:   domain.NewUserid(userid),
		Boardid:  domain.NewId(boardid),
		Priority: domain.NewPriority(priority),
		Contents: domain.NewContents(contents),
	}
}
