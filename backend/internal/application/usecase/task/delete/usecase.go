package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"
)

type TaskDeleteUsecaseInputPort interface {
	Delete(condition.TaskDeleteCondition) error
}
