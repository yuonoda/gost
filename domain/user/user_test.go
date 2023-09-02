package user

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"gost/domain"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	uuid1 := "e78a35ec-48ab-11ee-be56-0242ac120002"
	type args struct {
		idStr   string
		nameStr string
		role    int
	}
	tests := []struct {
		name      string
		args      args
		want      User
		wantError error
	}{
		{
			name: "pass",
			args: args{
				idStr:   uuid1,
				nameStr: "test",
				role:    1,
			},
			want: User{
				id:   id(uuid.MustParse(uuid1)),
				name: name("test"),
				role: admin,
			},
			wantError: nil,
		},
		{
			name: "fail - invalid format of uuid",
			args: args{
				idStr:   "xxx",
				nameStr: "test",
			},
			want:      User{},
			wantError: domain.InternalError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.idStr, tt.args.nameStr, tt.args.role)
			opts := cmp.Options{
				cmp.AllowUnexported(User{}),
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
