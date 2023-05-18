package model

import (
	"github.com/asaringo99/task_management/http/auth/entity"
	domain "github.com/asaringo99/task_management/internal/domain/entity"
)

type AuthModelInterface interface {
	ContainUser(user *entity.Userinfo) error
	CanAuthenticateUser(user *entity.Userinfo) error
	RegisterUser(*entity.Userinfo) error
	TransformUserIdFromUserName(*domain.Username) (domain.Userid, error)
}
