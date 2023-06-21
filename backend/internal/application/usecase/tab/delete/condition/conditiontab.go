package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type TabDeleteByTabCondition struct {
	Tabid  domain.Id
	Userid domain.Userid
}

func NewTabDeleteByTabCondition(tabid int, userid int) TabDeleteByTabCondition {
	return TabDeleteByTabCondition{
		Tabid:  domain.NewId(tabid),
		Userid: domain.NewUserid(userid),
	}
}
