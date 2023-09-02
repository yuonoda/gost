package user_test

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"gost/domain"
	"gost/domain/user"
	"testing"
)

func TestNew(t *testing.T) {
	uuid1 := "e78a35ec-48ab-11ee-be56-0242ac120002"
	type args struct {
		idStr   string
		nameStr string
	}
	tests := []struct {
		name      string
		args      args
		want      user.User
		wantError error
	}{
		{
			name: "pass",
			args: args{
				idStr:   uuid1,
				nameStr: "test",
			},
			want: (func() user.User {
				u, err := user.New(uuid1, "test")
				if err != nil {
					t.Fatal(err)
				}
				return u
			})(),
			wantError: nil,
		},
		{
			name: "fail - idStr is empty",
			args: args{
				idStr:   "",
				nameStr: "test",
			},
			want: user.User{},
			//want: (func() user.User {
			//	u, err := user.New("", "test")
			//	if err != nil {
			//		t.Fatal(err)
			//	}
			//	return u
			//})(),
			wantError: domain.InternalError{},
		},
		{
			name: "fail - invalid format of uuid",
			args: args{
				idStr:   "xxx",
				nameStr: "test",
			},
			want:      user.User{},
			wantError: domain.InternalError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := user.New(tt.args.idStr, tt.args.nameStr)
			opts := cmp.Options{
				cmp.AllowUnexported(user.User{}),
			}
			if diff := cmp.Diff(tt.want, got, opts); diff != "" {
				t.Errorf("diff: %v", diff)
			}
			if !errors.Is(err, tt.wantError) {
				t.Errorf("want error: %v, got: %v", tt.wantError, err)
			}
		})
	}
}
