package repository

import (
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type BoardCreateRepositoryInterface interface {
	Create(BoardCreateRepositoryInput) (*BoardCreateRepositoryOutput, error)
}

type BoardCreateRepositoryInput struct {
	Userid   domain.Userid
	Tabid    domain.Id
	Priority domain.Priority
	Status   domain.Status
}

type BoardCreateRepositoryOutput struct {
	Boardid  domain.Id
	Userid   domain.Userid
	Tabid    domain.Id
	Priority domain.Priority
	Status   domain.Status
}
