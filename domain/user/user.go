package user

import "fmt"

type User struct {
	id   id
	name name
	role role
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
func New(idStr string, nameStr string, roleInt int) (User, error) {
	newId, err := newID(idStr)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
	}
	r, err := newRole(roleInt)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return User{
		id:   newId,
		name: newName(nameStr),
		role: r,
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

func (u User) Role() int {
	return u.role.int()
}
