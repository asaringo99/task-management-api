package repository

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type TaskFetchRepositoryMock struct{}

func (repository TaskFetchRepositoryMock) Find() []TaskFetchRepositoryOutput {
	return []TaskFetchRepositoryOutput{
		{
			Taskid:   domain.NewTaskid(1),
			Userid:   domain.NewUserid(1),
			Boardid:  domain.NewId(1),
			Contents: domain.NewContents("test1"),
			Priority: domain.NewPriority(1),
		},
		{
			Taskid:   domain.NewTaskid(2),
			Userid:   domain.NewUserid(1),
			Boardid:  domain.NewId(2),
			Contents: domain.NewContents("test2"),
			Priority: domain.NewPriority(10),
		},
		{
			Taskid:  domain.NewTaskid(3),
			Userid:  domain.NewUserid(1),
			Boardid: domain.NewId(1),

			Contents: domain.NewContents("test3"),
			Priority: domain.NewPriority(5),
		},
	}
}
