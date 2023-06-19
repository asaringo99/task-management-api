package login

import (
	"fmt"

	"github.com/asaringo99/task_management/http/auth/entity"
	"github.com/asaringo99/task_management/http/auth/repository"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type LoginUsecase struct {
	repository repository.AuthRepositoryInterface
}

func (usecase *LoginUsecase) Login(userCredential entity.UserCredential) error {
	if err := usecase.canAuthenticate(userCredential); err != nil {
		return err
	}
	return nil
}

func (usecase *LoginUsecase) canAuthenticate(userCredential entity.UserCredential) error {
	fmt.Println(userCredential)
	username := domain.NewUsername(userCredential.Username)
	password := domain.NewPassword(userCredential.Password)
	user := entity.NewUserinfo(username, password)
	if err := usecase.repository.CanAuthenticateUser(&user); err != nil {
		return err
	}
	return nil
}

func NewLoginUsecase(repository repository.AuthRepositoryInterface) *LoginUsecase {
	return &LoginUsecase{
		repository: repository,
	}
}
