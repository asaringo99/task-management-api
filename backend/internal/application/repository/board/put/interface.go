package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/put/condition"
)

type BoardPutRepositoryInterface interface {
	Put(input condition.BoardPutCondition) error
}
