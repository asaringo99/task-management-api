package login

import (
	"github.com/asaringo99/task_management/http/auth/entity"
	"github.com/asaringo99/task_management/http/auth/usecase/login"
)

type LoginController struct {
	usecase login.LoginUsecase
}

func (controller *LoginController) Login(userCredential entity.UserCredential) error {
	if err := controller.usecase.Login(userCredential); err != nil {
		return err
	}
	return nil
}

func NewLoginController(usecase *login.LoginUsecase) *LoginController {
	return &LoginController{
		usecase: *usecase,
	}
}
