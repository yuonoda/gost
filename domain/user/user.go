package user

type User struct {
	id   id
	name name
}

func New(idStr string, nameStr string) (User, error) {
	// TODO ここでidStrが空文字の場合にuuidを生成するようにしたい
	newId, err := newID(idStr)
	if err != nil {
		return User{}, err
	}

	return User{
		id:   newId,
		name: newName(nameStr),
	}, nil
}

func (u User) Name() string {
	return u.name.string()
}

func (u User) SetID(s string) (User, error) {
	newId, err := newID(s)
	if err != nil {
		return User{}, err
	}
	u.id = newId
	return u, nil
}

func (u User) ID() string {
	return u.id.string()
}
