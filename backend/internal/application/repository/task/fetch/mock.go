package repository

import domain "github.com/asaringo99/task_management/internal/domain/entity"

type TaskFetchRepositoryMock struct{}

func (repository TaskFetchRepositoryMock) Find() []TaskFetchRepositoryOutput {
	return []TaskFetchRepositoryOutput{
		{
			Taskid:   domain.NewTaskid(1),
			Userid:   domain.NewUserid(1),
			Status:   domain.NewStatus("Pending"),
			Contents: domain.NewContents("test1"),
			Priority: domain.NewPriority(1),
		},
		{
			Taskid:   domain.NewTaskid(2),
			Userid:   domain.NewUserid(1),
			Status:   domain.NewStatus("In Progress"),
			Contents: domain.NewContents("test2"),
			Priority: domain.NewPriority(10),
		},
		{
			Taskid:   domain.NewTaskid(3),
			Userid:   domain.NewUserid(1),
			Status:   domain.NewStatus("Closed"),
			Contents: domain.NewContents("test3"),
			Priority: domain.NewPriority(5),
		},
	}
}