package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type TaskDeleteCondition struct {
	Userid domain.Userid
	Taskid domain.Taskid
}

func NewTaskDeleteCondition(userid int, taskid int) TaskDeleteCondition {
	return TaskDeleteCondition{
		Userid: domain.NewUserid(userid),
		Taskid: domain.NewTaskid(taskid),
	}
}
