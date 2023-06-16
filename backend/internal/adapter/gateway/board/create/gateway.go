package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	"gorm.io/gorm"
)

type BoardCreateGateway struct {
	db *gorm.DB
}

func (c *BoardCreateGateway) Create(input model.BoardModel) (model.BoardModel, error) {
	task := model.BoardModel{
		Userid:   input.Userid,
		Priority: input.Priority,
		Status:   input.Status,
	}
	c.db.Table("boards").Create(&task)
	return task, nil
}

func NewBoardCreateGateway(db *gorm.DB) BoardCreateGatewayInterface {
	return &BoardCreateGateway{
		db: db,
	}
}
