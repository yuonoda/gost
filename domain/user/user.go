package user

type User struct {
	id   id
	name name
}

func New(idStr string, nameStr string) User {
	var newId id
	// TODO ここでidStrが空文字の場合にuuidを生成するようにしたい
	if idStr != "" {
		newId = newID(idStr)
	}

	return User{
		id:   newId,
		name: newName(nameStr),
	}
}

func (u User) Name() string {
	return u.name.string()
}

func (u User) SetID(s string) User {
	u.id = newID(s)
	return u
}

func (u User) ID() string {
	return u.id.string()
}
