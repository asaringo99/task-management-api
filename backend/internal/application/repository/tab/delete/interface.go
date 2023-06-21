package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/delete/condition"
)

type TabDeleteRepositoryInterface interface {
	Delete(input condition.TabDeleteCondition) error
}
