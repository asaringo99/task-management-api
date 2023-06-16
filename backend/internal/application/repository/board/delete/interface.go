package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/delete/condition"
)

type BoardDeleteRepositoryInterface interface {
	Delete(input condition.BoardDeleteCondition) error
}
