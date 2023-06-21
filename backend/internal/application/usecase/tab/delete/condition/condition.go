package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type TabDeleteCondition struct {
	Tabid  domain.Id
	Userid domain.Userid
}

func NewTabDeleteCondition(tabid int, userid int) TabDeleteCondition {
	return TabDeleteCondition{
		Tabid:  domain.NewId(tabid),
		Userid: domain.NewUserid(userid),
	}
}
