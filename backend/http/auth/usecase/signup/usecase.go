package signup

import (
	"fmt"

	"github.com/asaringo99/task_management/http/auth/entity"
	"github.com/asaringo99/task_management/http/auth/repository"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
	"github.com/labstack/echo/v4"
)

type SignUpUsecase struct {
	repository repository.AuthRepositoryInterface
}

func (usecase *SignUpUsecase) RegisterUser(c echo.Context) error {
	userCredential := new(entity.UserCredential)
	if err := c.Bind(userCredential); err != nil {
		return err
	}
	fmt.Println(userCredential)
	username := domain.NewUsername(userCredential.Username)
	password := domain.NewHashedPassword(userCredential.Password)
	user := entity.NewUserinfo(username, password)
	if err := usecase.repository.ContainUser(&user); err != nil {
		return err
	}
	usecase.repository.Register(&user)
	return nil
}

func NewSignUpUsecase(repository repository.AuthRepositoryInterface) *SignUpUsecase {
	return &SignUpUsecase{
		repository: repository,
	}
}
