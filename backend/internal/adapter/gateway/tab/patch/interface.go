package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/tab/patch/condition"

type TabPatchGatewayInterface interface {
	Patch(input condition.TabPatchCondition) error
}
