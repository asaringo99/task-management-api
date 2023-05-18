package entity

import domain "github.com/asaringo99/task_management/internal/domain/entity"

type User struct {
	username domain.Username
	userid   domain.Userid
}

func NewUser(username domain.Username, userid domain.Userid) User {
	return User{
		username: username,
		userid:   userid,
	}
}

func (user *User) GetUsername() domain.Username {
	return user.username
}

func (user *User) GetUserid() domain.Userid {
	return user.userid
}
