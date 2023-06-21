package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/board/patch/condition"
)

type BoardPatchUsecaseInputPort interface {
	Patch(condition.BoardPatchCondition) error
}
