package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/task/put/condition"
)

type TaskPutRepositoryInterface interface {
	Put(input condition.TaskPutCondition) error
}
