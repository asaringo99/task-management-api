package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type BoardFetchCondition struct {
	Userid domain.Userid
}

func NewBoardFetchCondition(userid int) BoardFetchCondition {
	return BoardFetchCondition{
		Userid: domain.NewUserid(userid),
	}
}
