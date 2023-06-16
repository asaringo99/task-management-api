package entity

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestNewUserInfo(t *testing.T) {
	tests := []struct {
		name     string
		username domain.Username
		password domain.Password
		want     Userinfo
	}{
		{
			name:     "no diff",
			username: domain.NewUsername("user"),
			password: domain.NewPassword("passw0rd"),
			want: Userinfo{
				domain.NewUsername("user"),
				domain.NewPassword("passw0rd"),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewUserinfo(tc.username, tc.password)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
