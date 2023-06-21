package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type BoardDeleteByTabCondition struct {
	Tabid  domain.Id
	Userid domain.Userid
}

func NewBoardDeleteByTabCondition(tabid int, userid int) BoardDeleteByTabCondition {
	return BoardDeleteByTabCondition{
		Tabid:  domain.NewId(tabid),
		Userid: domain.NewUserid(userid),
	}
}
