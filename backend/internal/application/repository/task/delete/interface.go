package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/delete/condition"
)

type TaskDeleteRepositoryInterface interface {
	Delete(input condition.TaskDeleteCondition) error
}
