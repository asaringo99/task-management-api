package token

import (
	"github.com/asaringo99/task_management/http/auth/usecase/token"
	"github.com/labstack/echo/v4"
)

type TokenController struct {
	usecase token.TokenUsecase
}

func (controller *TokenController) RetrieveToken(c echo.Context) (string, error) {
	token, err := controller.usecase.RetrieveToken(c)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewTokenController(usecase *token.TokenUsecase) *TokenController {
	return &TokenController{
		usecase: *usecase,
	}
}
