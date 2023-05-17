package domain

type Contents struct {
	value string
}

func (d *Contents) ToValue() string {
	return d.value
}

func NewContents(value string) Contents {
	return Contents{
		value: value,
	}
}
