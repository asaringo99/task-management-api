package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	"github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"
	"gorm.io/gorm"
)

type TaskDeleteGateway struct {
	db *gorm.DB
}

func (c *TaskDeleteGateway) Delete(input condition.TaskDeleteCondition) error {
	var task []model.TaskModel
	c.db.Table("tasks").Where(&model.TaskModel{Id: input.Taskid.ToValue(), Userid: input.Userid.ToValue()}).Delete(task)
	return nil
}

func NewTaskDeleteGateway(db *gorm.DB) TaskDeleteGatewayInterface {
	return &TaskDeleteGateway{
		db: db,
	}
}
