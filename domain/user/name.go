package user

type name string

func (n name) string() string {
	return string(n)
}

func newName(str string) name {
	return name(str)
}
