package repository

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type BoardFetchRepositoryMock struct{}

func (repository BoardFetchRepositoryMock) Find() []BoardFetchRepositoryOutput {
	return []BoardFetchRepositoryOutput{
		{
			Boardid:  domain.NewId(1),
			Userid:   domain.NewUserid(1),
			Priority: domain.NewPriority(1),
			Status:   domain.NewStatus("test1"),
		},
		{
			Boardid:  domain.NewId(2),
			Userid:   domain.NewUserid(1),
			Priority: domain.NewPriority(10),
			Status:   domain.NewStatus("test2"),
		},
		{
			Boardid:  domain.NewId(3),
			Userid:   domain.NewUserid(1),
			Priority: domain.NewPriority(5),
			Status:   domain.NewStatus("test3"),
		},
	}
}
