package user

type User struct {
	id   id
	name name
}

func New(id uint64, name string) *User {
	return &User{
		id:   NewID(id),
		name: NewName(name),
	}
}
