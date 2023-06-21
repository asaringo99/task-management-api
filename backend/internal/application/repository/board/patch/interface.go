package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/patch/condition"
)

type BoardPatchRepositoryInterface interface {
	Patch(input condition.BoardPatchCondition) error
}
