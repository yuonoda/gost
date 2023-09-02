package user

import "gost/domain"

type role int

func (r role) int() int {
	return int(r)
}

const (
	normal role = iota
	admin  role = iota
)

func newRole(s int) (role, error) {
	switch s {
	case 0:
		return normal, nil
	case 1:
		return admin, nil
	default:
		return normal, domain.NewValidationError()
	}
}
