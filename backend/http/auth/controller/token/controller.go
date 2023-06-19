package token

import (
	"github.com/asaringo99/task_management/http/auth/entity"
	"github.com/asaringo99/task_management/http/auth/usecase/token"
)

type TokenController struct {
	usecase token.TokenUsecase
}

func (controller *TokenController) RetrieveToken(userCredential entity.UserCredential) (string, error) {
	token, err := controller.usecase.RetrieveToken(userCredential)
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
