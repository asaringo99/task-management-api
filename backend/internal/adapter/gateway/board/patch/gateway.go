package gateway

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/patch/condition"
	"gorm.io/gorm"
)

type BoardPatchGateway struct {
	db *gorm.DB
}

func (c *BoardPatchGateway) Patch(input condition.BoardPatchCondition) error {
	taskid := input.GetBoardid()
	userid := input.GetUserid()
	patch := input.GetPatchData()
	column := patch.Column
	value := patch.Value
	c.db.Table("tasks").Where("id = ? AND userid = ?", taskid.ToValue(), userid.ToValue()).Update(column, value)
	return nil
}

func NewBoardPatchGateway(db *gorm.DB) BoardPatchGatewayInterface {
	return &BoardPatchGateway{
		db: db,
	}
}
