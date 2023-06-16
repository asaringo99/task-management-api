package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/put/condition"
)

type BoardPutUsecaseInputPort interface {
	Put(condition.BoardPutCondition) error
}
