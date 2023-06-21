package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/delete/condition"
	"gorm.io/gorm"
)

type TabDeleteGateway struct {
	db *gorm.DB
}

func (c *TabDeleteGateway) Delete(input condition.TabDeleteCondition) error {
	var task []model.TabModel
	c.db.Table("tabs").Where(&model.TabModel{Id: input.Tabid.ToValue(), Userid: input.Userid.ToValue()}).Delete(task)
	return nil
}

func NewTabDeleteGateway(db *gorm.DB) TabDeleteGatewayInterface {
	return &TabDeleteGateway{
		db: db,
	}
}
