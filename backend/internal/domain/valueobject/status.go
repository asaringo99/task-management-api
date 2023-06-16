package domain

type Status struct {
	value string
}

func (d *Status) ToValue() string {
	return d.value
}

func NewStatus(value string) Status {
	return Status{
		value: value,
	}
}
