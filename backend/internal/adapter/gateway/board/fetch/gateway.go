package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	"github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"
	"gorm.io/gorm"
)

type taskFetchGateway struct {
	db *gorm.DB
}

func (c *taskFetchGateway) Find(input condition.BoardFetchCondition) ([]model.BoardModel, error) {
	var result []model.BoardModel
	c.db.Table("boards").Where(&model.BoardModel{Userid: input.Userid.ToValue()}).Find(&result)
	return result, nil
}

func NewBoardFetchGateway(db *gorm.DB) BoardFetchGatewayInterface {
	return &taskFetchGateway{
		db: db,
	}
}
