package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type BoardDeleteCondition struct {
	Boardid domain.Id
	Userid  domain.Userid
}

func NewBoardDeleteCondition(boardid int, userid int) BoardDeleteCondition {
	return BoardDeleteCondition{
		Boardid: domain.NewId(boardid),
		Userid:  domain.NewUserid(userid),
	}
}
