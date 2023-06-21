package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"
	"gorm.io/gorm"
)

type taskFetchGateway struct {
	db *gorm.DB
}

func (c *taskFetchGateway) Find(input condition.TabFetchCondition) ([]model.TabModel, error) {
	var result []model.TabModel
	c.db.Table("tabs").Where(&model.TabModel{Userid: input.Userid.ToValue()}).Find(&result)
	return result, nil
}

func NewTabFetchGateway(db *gorm.DB) TabFetchGatewayInterface {
	return &taskFetchGateway{
		db: db,
	}
}
