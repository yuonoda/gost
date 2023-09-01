package usecase

import "gost/domain/user"

// TODO moqとmockgenの比較

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	SaveUser(user user.User) (user.User, error)
}
