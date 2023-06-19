package token

import (
	"fmt"

	"github.com/asaringo99/task_management/http/auth/entity"
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	"github.com/asaringo99/task_management/http/auth/repository"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type TokenUsecase struct {
	repository repository.AuthRepositoryInterface
}

func (usecase *TokenUsecase) RetrieveToken(userCredential entity.UserCredential) (string, error) {
	username := domain.NewUsername(userCredential.Username)
	userid, err := usecase.repository.TransformUserIdFromUserName(&username)
	fmt.Println(userid, username)
	if err != nil {
		return "", err
	}
	token, err := authjwt.CreateJwtToken(userid.ToValue(), false)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewTokenUsecase(repository repository.AuthRepositoryInterface) *TokenUsecase {
	return &TokenUsecase{
		repository: repository,
	}
}
