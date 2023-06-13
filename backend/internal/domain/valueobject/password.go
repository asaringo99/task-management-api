package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func (d *Password) ToValue() string {
	return d.value
}

func NewPassword(value string) Password {
	return Password{
		value: value,
	}
}

func NewHashedPassword(value string) Password {
	password, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		panic("Not Permitted Password!")
	}
	return Password{
		value: string(password),
	}
}
