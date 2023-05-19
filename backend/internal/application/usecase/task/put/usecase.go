package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"
)

type TaskPutUsecaseInputPort interface {
	Put(condition.TaskPutCondition) error
}
