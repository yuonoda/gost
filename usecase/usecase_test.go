package usecase_test

import (
	"github.com/google/go-cmp/cmp"
	"gost/domain/user"
	"gost/usecase"
	"testing"
)

func TestUsecase_CreateUser(t *testing.T) {
	t.Parallel()
	type fields struct {
		userRepoFunc func(t *testing.T) usecase.Repository
	}
	type args struct {
		dto usecase.UserDTO
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   usecase.UserDTO
	}{
		{
			name: "pass",
			args: args{
				dto: usecase.UserDTO{
					Name: "test",
				},
			},
			fields: fields{
				userRepoFunc: func(t *testing.T) usecase.Repository {
					return &usecase.RepositoryMock{
						SaveUserFunc: func(userParam user.User) (user.User, error) {
							newUser := userParam.SetID(1)
							return newUser, nil
						},
					}
				},
			},
			want: usecase.UserDTO{
				ID:   1,
				Name: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.Usecase{
				Repo: tt.fields.userRepoFunc(t),
			}
			got, _ := u.CreateUser(tt.args.dto)
			// TODO エラー判定のベストプラクティス
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Usecase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
