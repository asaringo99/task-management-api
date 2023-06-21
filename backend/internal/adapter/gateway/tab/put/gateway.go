package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/put/condition"
	"gorm.io/gorm"
)

type TabPutGateway struct {
	db *gorm.DB
}

func (c *TabPutGateway) Put(input condition.TabPutCondition) error {
	userid := input.GetUserid()
	tabid := input.GetTabid()
	isActive := input.GetIsActive()
	title := input.GetTitle()
	c.db.Table("tabs").Where(&model.TabModel{Id: tabid.ToValue(), Userid: userid.ToValue()}).Updates(
		model.TabModel{
			Id:       tabid.ToValue(),
			Userid:   userid.ToValue(),
			IsActive: isActive,
			Title:    title.ToValue(),
		},
	)
	return nil
}

func NewTabPutGateway(db *gorm.DB) TabPutGatewayInterface {
	return &TabPutGateway{
		db: db,
	}
}
