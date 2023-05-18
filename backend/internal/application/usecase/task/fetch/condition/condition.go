package condition

import domain "github.com/asaringo99/task_management/internal/domain/entity"

type TaskFetchCondition struct {
	Userid domain.Userid
}

func NewTaskFetchCondition(userid int) TaskFetchCondition {
	return TaskFetchCondition{
		Userid: domain.NewUserid(userid),
	}
}
