package domain

type Id struct {
	value int
}

func (d *Id) ToValue() int {
	return d.value
}

func NewId(value int) Id {
	return Id{
		value: value,
	}
}
