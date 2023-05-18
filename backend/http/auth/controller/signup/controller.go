package signup

import (
	"github.com/asaringo99/task_management/http/auth/usecase/signup"
	"github.com/labstack/echo/v4"
)

type SignUpController struct {
	usecase signup.SignUpUsecase
}

func (controller *SignUpController) SignUp(c echo.Context) error {
	if err := controller.usecase.RegisterUser(c); err != nil {
		return err
	}
	return nil
}

func NewSignUpController(usecase *signup.SignUpUsecase) *SignUpController {
	return &SignUpController{
		usecase: *usecase,
	}
}
