package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/delete/condition"
)

type TabDeleteUsecaseInputPort interface {
	Delete(condition.TabDeleteCondition) error
}
