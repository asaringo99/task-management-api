package usecase

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestNewCreateUsecaseInput(t *testing.T) {
	tests := []struct {
		name     string
		tabid    int
		userid   int
		isActive bool
		title    string
		want     tabCreateUsecaseInput
	}{
		{
			name:     "no diff1",
			tabid:    1,
			userid:   1,
			isActive: true,
			title:    "test1",
			want: tabCreateUsecaseInput{
				Tabid:    domain.NewId(1),
				Userid:   domain.NewUserid(1),
				IsActive: true,
				Title:    domain.NewTitle("test1"),
			},
		},
		{
			name:     "no diff2",
			tabid:    1,
			userid:   1000,
			isActive: false,
			title:    "test2",
			want: tabCreateUsecaseInput{
				Tabid:    domain.NewId(1),
				Userid:   domain.NewUserid(1000),
				IsActive: false,
				Title:    domain.NewTitle("test2"),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewTabCreateUsecaseInput(tc.userid, tc.tabid, tc.isActive, tc.title)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
