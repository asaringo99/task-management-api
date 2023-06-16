package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	"gorm.io/gorm"
)

type TaskCreateGateway struct {
	db *gorm.DB
}

func (c *TaskCreateGateway) Create(input model.TaskModel) (model.TaskModel, error) {
	task := model.TaskModel{
		Userid:   input.Userid,
		Boardid:  input.Boardid,
		Priority: input.Priority,
		Contents: input.Contents,
	}
	c.db.Table("tasks").Create(&task)
	return task, nil
}

func NewTaskCreateGateway(db *gorm.DB) TaskCreateGatewayInterface {
	return &TaskCreateGateway{db: db}
}
