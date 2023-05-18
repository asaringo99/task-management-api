package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
	"gorm.io/gorm"
)

type taskFetchGateway struct {
	db *gorm.DB
}

func (c *taskFetchGateway) Find(input condition.TaskFetchCondition) ([]model.TaskModel, error) {
	var result []model.TaskModel
	c.db.Table("tasks").Where(&model.TaskModel{Userid: input.Userid.ToValue()}).Find(&result)
	return result, nil
}

func NewTaskFetchGateway(db *gorm.DB) TaskFetchGatewayInterface {
	return &taskFetchGateway{
		db: db,
	}
}
