package condition

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestNewBoardPutCondition(t *testing.T) {
	tests := []struct {
		name     string
		boardid  int
		userid   int
		priority int
		status   string
		want     BoardPutCondition
	}{
		{
			name:     "no diff1",
			userid:   1,
			boardid:  1,
			status:   "test",
			priority: 1,
			want: BoardPutCondition{
				boardid:  domain.NewId(1),
				userid:   domain.NewUserid(1),
				status:   domain.NewStatus("test"),
				priority: domain.NewPriority(1),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewBoardPutCondition(tc.boardid, tc.userid, tc.priority, tc.status)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
