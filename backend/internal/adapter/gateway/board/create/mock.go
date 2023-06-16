package gateway

import model "github.com/asaringo99/task_management/internal/adapter/gateway/board"

type BoardAllGetGatewayMock struct{}

func (c *BoardAllGetGatewayMock) Create(input model.BoardModel) (model.BoardModel, error) {
	return model.BoardModel{}, nil
}

func NewBoardAllGetGatewayMock() BoardCreateGatewayInterface {
	return &BoardAllGetGatewayMock{}
}
