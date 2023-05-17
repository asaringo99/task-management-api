package domain

type Taskid struct {
	value int
}

func (d *Taskid) ToValue() int {
	return d.value
}

func NewTaskid(value int) Taskid {
	return Taskid{
		value: value,
	}
}
