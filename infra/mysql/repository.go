package mysql

import "gost/domain/user"

type Repository struct{}

func (r Repository) SaveUser(userParam user.User) (user.User, error) {
	return user.User{}, nil
}
