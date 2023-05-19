package model

import (
	"fmt"

	"github.com/asaringo99/task_management/http/auth/entity"
	domain "github.com/asaringo99/task_management/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userInfoModel struct {
	Username string
	Password string
}

type userModel struct {
	Id       int
	Username string
}

type AuthModelImpl struct {
	db *gorm.DB
}

func (c *AuthModelImpl) ContainUser(user *entity.Userinfo) error {
	var usermodel []userInfoModel
	username := user.GetUsername()
	c.db.Table("users").Where(&userInfoModel{Username: username.ToValue()}).Find(&usermodel)
	if len(usermodel) == 0 {
		return nil
	}
	return fmt.Errorf("ContainUser")
}

func (c *AuthModelImpl) CanAuthenticateUser(user *entity.Userinfo) error {
	var usermodel userInfoModel
	username := user.GetUsername()
	password := user.GetPassword()
	c.db.Table("users").Where(&userInfoModel{Username: username.ToValue()}).Find(&usermodel)
	if err := bcrypt.CompareHashAndPassword([]byte(usermodel.Password), []byte(password.ToValue())); err != nil {
		return err
	}
	return nil
}

func (c *AuthModelImpl) RegisterUser(user *entity.Userinfo) error {
	username := user.GetUsername()
	password := user.GetPassword()
	registeredUser := userInfoModel{
		Username: username.ToValue(),
		Password: password.ToValue(),
	}
	fmt.Println(registeredUser)
	c.db.Table("users").Create(&registeredUser)
	return nil
}

func (c *AuthModelImpl) TransformUserIdFromUserName(u *domain.Username) (domain.Userid, error) {
	var usermodel userModel
	c.db.Table("users").Where(&userModel{Username: u.ToValue()}).First(&usermodel)
	userid := domain.NewUserid(usermodel.Id)
	return userid, nil
}

func NewAuthModel(db *gorm.DB) AuthModelInterface {
	return &AuthModelImpl{
		db: db,
	}
}
