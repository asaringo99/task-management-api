package domain

type Username struct {
	value string
}

func (d *Username) ToValue() string {
	return d.value
}

func NewUsername(value string) Username {
	return Username{
		value: value,
	}
}

type Userid struct {
	value int
}

func (d *Userid) ToValue() int {
	return d.value
}

func NewUserid(value int) Userid {
	return Userid{
		value: value,
	}
}
