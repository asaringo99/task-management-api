package repository

import (
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TabCreateRepositoryInterface interface {
	Create(TabCreateRepositoryInput) (*TabCreateRepositoryOutput, error)
}

type TabCreateRepositoryInput struct {
	Tabid    domain.Id
	Userid   domain.Userid
	IsActive bool
	Title    domain.Title
}

type TabCreateRepositoryOutput struct {
	Tabid    domain.Id
	Userid   domain.Userid
	IsActive bool
	Title    domain.Title
}
