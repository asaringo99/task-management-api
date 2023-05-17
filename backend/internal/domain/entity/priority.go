package domain

type Priority struct {
	value int
}

func (d *Priority) ToValue() int {
	return d.value
}

func NewPriority(value int) Priority {
	return Priority{
		value: value,
	}
}
