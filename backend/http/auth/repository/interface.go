package repository

import (
	"github.com/asaringo99/task_management/http/auth/entity"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type AuthRepositoryInterface interface {
	ContainUser(*entity.Userinfo) error
	CanAuthenticateUser(*entity.Userinfo) error
	Register(*entity.Userinfo) error
	TransformUserIdFromUserName(*domain.Username) (*domain.Userid, error)
}
