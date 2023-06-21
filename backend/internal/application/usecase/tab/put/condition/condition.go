package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type TabPutCondition struct {
	tabid    domain.Id
	userid   domain.Userid
	isActive bool
	title    domain.Title
}

func NewTabPutCondition(tabid int, userid int, isActive bool, title string) TabPutCondition {
	return TabPutCondition{
		tabid:    domain.NewId(tabid),
		userid:   domain.NewUserid(userid),
		isActive: isActive,
		title:    domain.NewTitle(title),
	}
}

func (c TabPutCondition) GetTabid() domain.Id {
	return c.tabid
}

func (c TabPutCondition) GetUserid() domain.Userid {
	return c.userid
}

func (c TabPutCondition) GetIsActive() bool {
	return c.isActive
}

func (c TabPutCondition) GetTitle() domain.Title {
	return c.title
}
