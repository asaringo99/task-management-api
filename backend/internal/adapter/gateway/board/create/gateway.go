package gateway

import (
	"fmt"

	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	"gorm.io/gorm"
)

type BoardCreateGateway struct {
	db *gorm.DB
}

func (c *BoardCreateGateway) Create(input model.BoardModel) (model.BoardModel, error) {
	task := model.BoardModel{
		Id:       input.Id,
		Userid:   input.Userid,
		Tabid:    input.Tabid,
		Priority: input.Priority,
		Status:   input.Status,
	}
	fmt.Println(task)
	c.db.Table("boards").Create(&task)
	return task, nil
}

func NewBoardCreateGateway(db *gorm.DB) BoardCreateGatewayInterface {
	return &BoardCreateGateway{
		db: db,
	}
}
