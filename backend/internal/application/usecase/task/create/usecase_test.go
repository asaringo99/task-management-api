package usecase

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestNewCreateUsecaseInput(t *testing.T) {
	tests := []struct {
		name     string
		userid   int
		boardid  int
		priority int
		contents string
		want     taskCreateUsecaseInput
	}{
		{
			name:     "no diff1",
			userid:   1,
			boardid:  1,
			priority: 1,
			contents: "this is the first test",
			want: taskCreateUsecaseInput{
				Userid:   domain.NewUserid(1),
				Boardid:  domain.NewId(1),
				Priority: domain.NewPriority(1),
				Contents: domain.NewContents("this is the first test"),
			},
		},
		{
			name:     "no diff2",
			userid:   1000,
			boardid:  2,
			priority: 20,
			contents: "this is the second test",
			want: taskCreateUsecaseInput{
				Userid:   domain.NewUserid(1000),
				Boardid:  domain.NewId(2),
				Priority: domain.NewPriority(20),
				Contents: domain.NewContents("this is the second test"),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewTaskCreateUsecaseInput(tc.userid, tc.boardid, tc.priority, tc.contents)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
