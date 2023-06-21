package condition

import domain "github.com/asaringo99/task_management/internal/domain/valueobject"

type PatchData struct {
	Column string
	Value  interface{}
}

type BoardPatchCondition struct {
	boardid domain.Id
	userid  domain.Userid
	fix     PatchData
}

func NewBoardPatchCondition(boardid int, userid int, patchData PatchData) BoardPatchCondition {
	return BoardPatchCondition{
		boardid: domain.NewId(boardid),
		userid:  domain.NewUserid(userid),
		fix:     patchData,
	}
}

func (c BoardPatchCondition) GetBoardid() domain.Id {
	return c.boardid
}

func (c BoardPatchCondition) GetUserid() domain.Userid {
	return c.userid
}

func (c BoardPatchCondition) GetPatchData() PatchData {
	return c.fix
}
