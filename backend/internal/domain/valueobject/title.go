package domain

type Title struct {
	value string
}

func (d *Title) ToValue() string {
	return d.value
}

func NewTitle(value string) Title {
	return Title{
		value: value,
	}
}
