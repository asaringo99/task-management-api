package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	"github.com/asaringo99/task_management/internal/application/usecase/board/put/condition"
	"gorm.io/gorm"
)

type BoardPutGateway struct {
	db *gorm.DB
}

func (c *BoardPutGateway) Put(input condition.BoardPutCondition) error {
	userid := input.GetUserid()
	boardid := input.GetBoardid()
	priority := input.GetPriority()
	status := input.GetStatus()
	c.db.Table("boards").Where(&model.BoardModel{Id: boardid.ToValue(), Userid: userid.ToValue()}).Updates(
		model.BoardModel{
			Id:       boardid.ToValue(),
			Priority: priority.ToValue(),
			Status:   status.ToValue(),
		},
	)
	return nil
}

func NewBoardPutGateway(db *gorm.DB) BoardPutGatewayInterface {
	return &BoardPutGateway{
		db: db,
	}
}
