package dto

import (
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

type UserNamePresenterDto struct {
	value string
	unit  string
}

func (d *UserNamePresenterDto) ToValue() string {
	return d.value
}

func (d *UserNamePresenterDto) ToUnit() string {
	return d.unit
}

func NewUserNamePresenterDto(username domain.Username) *UserNamePresenterDto {
	return &UserNamePresenterDto{
		value: username.ToValue(),
	}
}
