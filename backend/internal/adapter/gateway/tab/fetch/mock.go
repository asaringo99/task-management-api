package gateway

import (
	model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"
	"github.com/asaringo99/task_management/internal/application/usecase/tab/fetch/condition"
)

type TabAllGetGatewayMock struct{}

func (c *TabAllGetGatewayMock) Find(input condition.TabFetchCondition) ([]model.TabModel, error) {
	return []model.TabModel{
		{
			Id:       1,
			Userid:   1,
			IsActive: false,
			Title:    "test1",
		},
		{
			Id:       2,
			Userid:   1,
			IsActive: false,
			Title:    "test2",
		},
		{
			Id:       3,
			Userid:   2,
			IsActive: true,
			Title:    "test3",
		},
	}, nil
}

func NewTabAllGetGatewayMock() TabFetchGatewayInterface {
	return &TabAllGetGatewayMock{}
}
