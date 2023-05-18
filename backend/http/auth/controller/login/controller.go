package login

import (
	"github.com/asaringo99/task_management/http/auth/usecase/login"
	"github.com/labstack/echo/v4"
)

type LoginController struct {
	usecase login.LoginUsecase
}

func (controller *LoginController) Login(c echo.Context) error {
	if err := controller.usecase.Login(c); err != nil {
		return err
	}
	return nil
}

func NewLoginController(usecase *login.LoginUsecase) *LoginController {
	return &LoginController{
		usecase: *usecase,
	}
}
