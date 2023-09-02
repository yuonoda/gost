package user

import (
	"github.com/google/uuid"
	"gost/domain"
)

type id uuid.UUID

func (id id) string() string {
	return uuid.UUID(id).String()
}

func newID(str string) (id, error) {
	parsed, err := uuid.Parse(str)
	if err != nil {
		return id(uuid.Nil), domain.NewInternalError()
	}
	return id(parsed), nil
}

func genID() (id, error) {
	generated, err := uuid.NewUUID()
	if err != nil {
		return id(uuid.Nil), domain.NewInternalError()
	}
	return id(generated), nil
}
