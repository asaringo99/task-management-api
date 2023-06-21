package gateway

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/patch/condition"
	"gorm.io/gorm"
)

type TabPatchGateway struct {
	db *gorm.DB
}

func (c *TabPatchGateway) Patch(input condition.TabPatchCondition) error {
	tabid := input.GetTabid()
	userid := input.GetUserid()
	patch := input.GetPatchData()
	column := patch.Column
	value := patch.Value
	c.db.Table("tasks").Where("id = ? AND userid = ?", tabid.ToValue(), userid.ToValue()).Update(column, value)
	return nil
}

func NewTabPatchGateway(db *gorm.DB) TabPatchGatewayInterface {
	return &TabPatchGateway{
		db: db,
	}
}
