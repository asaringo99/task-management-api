package repository

import (
	"github.com/asaringo99/task_management/internal/application/usecase/tab/put/condition"
)

type TabPutRepositoryInterface interface {
	Put(input condition.TabPutCondition) error
}
