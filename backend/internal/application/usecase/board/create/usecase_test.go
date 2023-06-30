package usecase

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestNewCreateUsecaseInput(t *testing.T) {
	tests := []struct {
		name     string
		userid   domain.Userid
		tabid    domain.Id
		priority domain.Priority
		status   domain.Status
		want     boardCreateUsecaseInput
	}{
		{
			name:     "no diff1",
			userid:   domain.NewUserid(1),
			tabid:    domain.NewId(1),
			priority: domain.NewPriority(1),
			status:   domain.NewStatus("Pending"),
			want: boardCreateUsecaseInput{
				Userid:   domain.NewUserid(1),
				Tabid:    domain.NewId(1),
				Priority: domain.NewPriority(1),
				Status:   domain.NewStatus("Pending"),
			},
		},
		{
			name:     "no diff2",
			userid:   domain.NewUserid(1000),
			tabid:    domain.NewId(1),
			priority: domain.NewPriority(20),
			status:   domain.NewStatus("Todo"),
			want: boardCreateUsecaseInput{
				Userid:   domain.NewUserid(1000),
				Tabid:    domain.NewId(1),
				Priority: domain.NewPriority(20),
				Status:   domain.NewStatus("Todo"),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewBoardCreateUsecaseInput(tc.userid, tc.tabid, tc.priority, tc.status)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
