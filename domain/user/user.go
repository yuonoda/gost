package user

import "fmt"

type User struct {
	id   id
	name name
}

// Create brand-new user with generated ID
func Create(nameStr string) (User, error) {
	newId, err := genID()
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return User{
		id:   newId,
		name: newName(nameStr),
	}, nil
}

// New is constructor of User
func New(idStr string, nameStr string) (User, error) {
	newId, err := newID(idStr)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
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
