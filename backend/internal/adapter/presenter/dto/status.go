package dto

import (
	domain "github.com/asaringo99/task_management/internal/domain/entity"
)

type StatusPresenterDto struct {
	value string
	unit  string
}

func (d *StatusPresenterDto) ToValue() string {
	return d.value
}

func (d *StatusPresenterDto) ToUnit() string {
	return d.unit
}

func NewStatusPresenterDto(status domain.Status) *StatusPresenterDto {
	return &StatusPresenterDto{
		value: status.ToValue(),
	}
}
