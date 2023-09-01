package user

import (
	"github.com/google/uuid"
)

type id uuid.UUID

func (id id) string() string {
	return uuid.UUID(id).String()
}

func newID(str string) id {
	return id(uuid.MustParse(str))
}

func genID() (id, error) {
	generated, err := uuid.NewUUID()
	if err != nil {
		return id(uuid.Nil), err
	}
	return id(generated), nil
}
