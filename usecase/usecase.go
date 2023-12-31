package usecase

import "gost/domain/user"

type Usecase struct {
	Repo Repository
}

func (uc *Usecase) CreateUser(dto UserDTO) (UserDTO, error) {
	u, err := user.Create(dto.Name)
	if err != nil {
		return UserDTO{}, err // TODO ユースケースエラーへの変換？
	}
	saved, err := uc.Repo.SaveUser(u)
	if err != nil {
		return UserDTO{}, err
	}
	return UserDTO{
		ID:   saved.ID(),
		Name: saved.Name(),
		Role: saved.Role(),
	}, nil
}
