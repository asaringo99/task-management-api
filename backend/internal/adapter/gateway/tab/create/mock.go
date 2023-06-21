package gateway

import model "github.com/asaringo99/task_management/internal/adapter/gateway/tab"

type TabAllGetGatewayMock struct{}

func (c *TabAllGetGatewayMock) Create(input model.TabModel) (model.TabModel, error) {
	return model.TabModel{}, nil
}

func NewTabAllGetGatewayMock() TabCreateGatewayInterface {
	return &TabAllGetGatewayMock{}
}
