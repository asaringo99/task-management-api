package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type PatchData struct {
	Column string
	Value  interface{}
}

type TabPatchCondition struct {
	tabid  domain.Id
	userid domain.Userid
	fix    PatchData
}

func NewTabPatchCondition(tabid int, userid int, patchData PatchData) TabPatchCondition {
	return TabPatchCondition{
		tabid:  domain.NewId(tabid),
		userid: domain.NewUserid(userid),
		fix:    patchData,
	}
}

func (c TabPatchCondition) GetTabid() domain.Id {
	return c.tabid
}

func (c TabPatchCondition) GetUserid() domain.Userid {
	return c.userid
}

func (c TabPatchCondition) GetPatchData() PatchData {
	return c.fix
}
