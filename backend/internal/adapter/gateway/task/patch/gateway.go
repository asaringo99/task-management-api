package gateway

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/patch/condition"
	"gorm.io/gorm"
)

type TaskPatchGateway struct {
	db *gorm.DB
}

func (c *TaskPatchGateway) Patch(input condition.TaskPatchCondition) error {
	taskid := input.GetTaskid()
	userid := input.GetUserid()
	patch := input.GetPatchData()
	column := patch.Column
	value := patch.Value
	c.db.Table("tasks").Where("id = ? AND userid = ?", taskid.ToValue(), userid.ToValue()).Update(column, value)
	return nil
}

func NewTaskPatchGateway(db *gorm.DB) TaskPatchGatewayInterface {
	return &TaskPatchGateway{
		db: db,
	}
}
