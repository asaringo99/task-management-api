package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/patch/condition"
)

type TabPatchUsecaseInputPort interface {
	Patch(condition.TabPatchCondition) error
}
