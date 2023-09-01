package user

type name string

func NewName(str string) name {
	return name(str)
}
