package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/patch/condition"
)

type TaskPatchRepositoryInterface interface {
	Patch(input condition.TaskPatchCondition) error
}
