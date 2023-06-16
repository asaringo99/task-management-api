package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/board"
	"github.com/asaringo99/task_management/internal/application/usecase/board/fetch/condition"
)

type BoardAllGetGatewayMock struct{}

func (c *BoardAllGetGatewayMock) Find(input condition.BoardFetchCondition) ([]model.BoardModel, error) {
	return []model.BoardModel{
		{
			Id:       1,
			Userid:   1,
			Priority: 1,
			Status:   "Pending",
		},
		{
			Id:       2,
			Userid:   1,
			Priority: 2,
			Status:   "Todo",
		},
		{
			Id:       3,
			Userid:   1,
			Priority: 3,
			Status:   "In Progress",
		},
	}, nil
}

func NewBoardAllGetGatewayMock() BoardFetchGatewayInterface {
	return &BoardAllGetGatewayMock{}
}
