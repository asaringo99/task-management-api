package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/task"
	"github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"
	"gorm.io/gorm"
)

type TaskPutGateway struct {
	db *gorm.DB
}

func (c *TaskPutGateway) Put(input condition.TaskPutCondition) error {
	taskid := input.GetTaskid()
	userid := input.GetUserid()
	boardid := input.GetBoardid()
	priority := input.GetPriority()
	contents := input.GetContents()
	c.db.Table("tasks").Where(&model.TaskModel{Id: taskid.ToValue(), Userid: userid.ToValue()}).Updates(
		model.TaskModel{
			Boardid:  boardid.ToValue(),
			Priority: priority.ToValue(),
			Contents: contents.ToValue(),
		},
	)
	return nil
}

func NewTaskPutGateway(db *gorm.DB) TaskPutGatewayInterface {
	return &TaskPutGateway{
		db: db,
	}
}
