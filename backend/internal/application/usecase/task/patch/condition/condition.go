package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type PatchData struct {
	Column string
	Value  int
}

type TaskPatchCondition struct {
	taskid domain.Taskid
	userid domain.Userid
	fix    PatchData
}

func NewTaskPatchCondition(taskid int, userid int, patchData PatchData) TaskPatchCondition {
	return TaskPatchCondition{
		taskid: domain.NewTaskid(taskid),
		userid: domain.NewUserid(userid),
		fix:    patchData,
	}
}

func (c TaskPatchCondition) GetTaskid() domain.Taskid {
	return c.taskid
}

func (c TaskPatchCondition) GetUserid() domain.Userid {
	return c.userid
}

func (c TaskPatchCondition) GetPatchData() PatchData {
	return c.fix
}
