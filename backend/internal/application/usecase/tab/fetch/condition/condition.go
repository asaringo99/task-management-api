package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type TabFetchCondition struct {
	Userid domain.Userid
}

func NewTabFetchCondition(userid int) TabFetchCondition {
	return TabFetchCondition{
		Userid: domain.NewUserid(userid),
	}
}
