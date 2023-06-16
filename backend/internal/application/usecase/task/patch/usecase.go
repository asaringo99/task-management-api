package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/patch/condition"
)

type TaskPatchUsecaseInputPort interface {
	Patch(condition.TaskPatchCondition) error
}
