package user_test

import (
	"github.com/google/go-cmp/cmp"
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
		name string
		args args
		want user.User
	}{
		{
			name: "pass",
			args: args{
				idStr:   uuid1,
				nameStr: "test",
			},
			want: user.New(uuid1, "test"),
		},
		{
			name: "pass - idStr is empty",
			args: args{
				idStr:   "",
				nameStr: "test",
			},
			want: user.New("", "test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := user.New(tt.args.idStr, tt.args.nameStr)
			opts := cmp.Options{
				cmp.AllowUnexported(user.User{}),
			}
			if diff := cmp.Diff(tt.want, got, opts); diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}
