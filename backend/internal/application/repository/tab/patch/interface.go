package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/patch/condition"
)

type TabPatchRepositoryInterface interface {
	Patch(input condition.TabPatchCondition) error
}
