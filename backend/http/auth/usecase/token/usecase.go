package token

import (
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	"github.com/asaringo99/task_management/http/auth/repository"
	domain "github.com/asaringo99/task_management/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type TokenUsecase struct {
	repository repository.AuthRepositoryInterface
}

func (usecase *TokenUsecase) RetrieveToken(c echo.Context) (string, error) {
	u := c.FormValue("username")
	username := domain.NewUsername(u)
	userid, err := usecase.repository.TransformUserIdFromUserName(&username)
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
