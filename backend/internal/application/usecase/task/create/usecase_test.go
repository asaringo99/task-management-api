package usecase

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/entity"
)

func TestNewCreateUsecaseInput(t *testing.T) {
	tests := []struct {
		name     string
		userid   int
		status   string
		priority int
		contents string
		want     taskCreateUsecaseInput
	}{
		{
			name:     "no diff1",
			userid:   1,
			status:   "In Progress",
			priority: 1,
			contents: "this is the first test",
			want: taskCreateUsecaseInput{
				Userid:   domain.NewUserid(1),
				Status:   domain.NewStatus("In Progress"),
				Priority: domain.NewPriority(1),
				Contents: domain.NewContents("this is the first test"),
			},
		},
		{
			name:     "no diff2",
			userid:   1000,
			status:   "Closed",
			priority: 20,
			contents: "this is the second test",
			want: taskCreateUsecaseInput{
				Userid:   domain.NewUserid(1000),
				Status:   domain.NewStatus("Closed"),
				Priority: domain.NewPriority(20),
				Contents: domain.NewContents("this is the second test"),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewTaskCreateUsecaseInput(tc.userid, tc.status, tc.priority, tc.contents)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
