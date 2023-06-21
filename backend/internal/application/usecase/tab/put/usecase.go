package usecase

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/put/condition"
)

type TabPutUsecaseInputPort interface {
	Put(condition.TabPutCondition) error
}
