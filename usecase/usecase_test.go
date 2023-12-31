package usecase_test

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
		name    string
		fields  fields
		args    args
		want    usecase.UserDTO
		wantErr error
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
							return userParam, nil
						},
					}
				},
			},
			want: usecase.UserDTO{
				ID:   "xxx",
				Name: "test",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.Usecase{
				Repo: tt.fields.userRepoFunc(t),
			}
			got, err := u.CreateUser(tt.args.dto)
			opts := cmp.Options{
				cmpopts.IgnoreFields(usecase.UserDTO{}, "ID"),
			}
			if diff := cmp.Diff(got, tt.want, opts); diff != "" {
				t.Errorf("Usecase.CreateUser() = %v, want %v", got, tt.want)
			}
			if !errors.Is(tt.wantErr, err) {
				t.Errorf("Usecase.CreateUser() error = %v, wantErr %v", got, tt.want)
			}
		})
	}
}
