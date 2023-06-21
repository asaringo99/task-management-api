package repository

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type TabFetchRepositoryMock struct{}

func (repository TabFetchRepositoryMock) Find() []TabFetchRepositoryOutput {
	return []TabFetchRepositoryOutput{
		{
			Tabid:    domain.NewId(1),
			Userid:   domain.NewUserid(1),
			IsActive: false,
			Title:    domain.NewTitle("test1"),
		},
		{
			Tabid:    domain.NewId(2),
			Userid:   domain.NewUserid(1),
			IsActive: false,
			Title:    domain.NewTitle("test2"),
		},
		{
			Tabid:    domain.NewId(3),
			Userid:   domain.NewUserid(1),
			IsActive: true,
			Title:    domain.NewTitle("test3"),
		},
	}
}
