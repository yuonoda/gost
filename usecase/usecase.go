package usecase

import "gost/domain/user"

type Usecase struct {
	Repo Repository
}

func (uc *Usecase) CreateUser(dto UserDTO) (UserDTO, error) {
	u := user.New(0, dto.Name)
	saved, err := uc.Repo.SaveUser(u)
	if err != nil {
		return UserDTO{}, err
	}
	return UserDTO{
		ID:   saved.ID(),
		Name: saved.Name(),
	}, nil
}
