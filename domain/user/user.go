package user

type User struct {
	id   id
	name name
}

func New(id uint64, name string) User {
	return User{
		id:   NewID(id),
		name: NewName(name),
	}
}

func (u User) Name() string {
	return u.name.String()
}

func (u User) SetID(i uint64) User {
	u.id = NewID(i)
	return u
}

func (u User) ID() uint64 {
	return uint64(u.id)
}
