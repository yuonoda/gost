package user

type id uint64

func NewID(i uint64) id {
	return id(i)
}
