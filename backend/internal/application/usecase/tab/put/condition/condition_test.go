package condition

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestNewTabPutCondition(t *testing.T) {
	tests := []struct {
		name     string
		tabid    int
		userid   int
		isActive bool
		title    string
		want     TabPutCondition
	}{
		{
			name:     "no diff1",
			userid:   1,
			tabid:    1,
			isActive: true,
			title:    "test",
			want: TabPutCondition{
				tabid:    domain.NewId(1),
				userid:   domain.NewUserid(1),
				isActive: true,
				title:    domain.NewTitle("test"),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewTabPutCondition(tc.tabid, tc.userid, tc.isActive, tc.title)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
