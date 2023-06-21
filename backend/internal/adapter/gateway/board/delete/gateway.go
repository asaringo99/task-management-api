package gateway

import (
	"fmt"

	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	"github.com/asaringo99/task_management/internal/application/usecase/board/delete/condition"
	"gorm.io/gorm"
)

type BoardDeleteGateway struct {
	db *gorm.DB
}

func (c *BoardDeleteGateway) Delete(input condition.BoardDeleteCondition) error {
	var task []model.BoardModel
	fmt.Println(input)
	c.db.Table("boards").Where(&model.BoardModel{Id: input.Boardid.ToValue(), Userid: input.Userid.ToValue()}).Delete(task)
	return nil
}

func NewBoardDeleteGateway(db *gorm.DB) BoardDeleteGatewayInterface {
	return &BoardDeleteGateway{
		db: db,
	}
}
