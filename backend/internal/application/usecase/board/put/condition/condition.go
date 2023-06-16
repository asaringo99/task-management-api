package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type BoardPutCondition struct {
	boardid  domain.Id
	userid   domain.Userid
	priority domain.Priority
	status   domain.Status
}

func NewBoardPutCondition(boardid int, userid int, priority int, status string) BoardPutCondition {
	return BoardPutCondition{
		boardid:  domain.NewId(boardid),
		userid:   domain.NewUserid(userid),
		priority: domain.NewPriority(priority),
		status:   domain.NewStatus(status),
	}
}

func (c BoardPutCondition) GetBoardid() domain.Id {
	return c.boardid
}

func (c BoardPutCondition) GetUserid() domain.Userid {
	return c.userid
}

func (c BoardPutCondition) GetPriority() domain.Priority {
	return c.priority
}

func (c BoardPutCondition) GetStatus() domain.Status {
	return c.status
}
