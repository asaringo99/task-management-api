package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type TaskPutCondition struct {
	taskid   domain.Taskid
	userid   domain.Userid
	boardid  domain.Id
	priority domain.Priority
	contents domain.Contents
}

func NewTaskPutCondition(taskid int, userid int, boardid int, priority int, contents string) TaskPutCondition {
	return TaskPutCondition{
		taskid:   domain.NewTaskid(taskid),
		userid:   domain.NewUserid(userid),
		boardid:  domain.NewId(boardid),
		priority: domain.NewPriority(priority),
		contents: domain.NewContents(contents),
	}
}

func (c TaskPutCondition) GetTaskid() domain.Taskid {
	return c.taskid
}

func (c TaskPutCondition) GetUserid() domain.Userid {
	return c.userid
}

func (c TaskPutCondition) GetBoardid() domain.Id {
	return c.boardid
}

func (c TaskPutCondition) GetPriority() domain.Priority {
	return c.priority
}

func (c TaskPutCondition) GetContents() domain.Contents {
	return c.contents
}
