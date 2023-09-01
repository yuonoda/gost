package user

type name string

func (n name) String() string {
	return string(n)
}

func NewName(str string) name {
	return name(str)
}
