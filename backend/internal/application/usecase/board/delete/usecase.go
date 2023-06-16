package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/delete/condition"
)

type BoardDeleteUsecaseInputPort interface {
	Delete(condition.BoardDeleteCondition) error
}
