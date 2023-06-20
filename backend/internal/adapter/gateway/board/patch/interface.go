package gateway

import "github.com/asaringo99/task_management/internal/application/usecase/board/patch/condition"

type BoardPatchGatewayInterface interface {
	Patch(input condition.BoardPatchCondition) error
}
