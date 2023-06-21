package usecase

import (
	"reflect"
	"testing"

	repository "github.com/asaringo99/task_management/internal/application/repository/board/fetch"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name   string
		output repository.BoardFetchRepositoryOutput
		want   BoardFetchUsecaseOutput
	}{
		{
			name: "no diff1",
			output: repository.BoardFetchRepositoryOutput{
				Boardid:  domain.NewId(1),
				Userid:   domain.NewUserid(1),
				Tabid:    domain.NewId(2),
				Priority: domain.NewPriority(1),
				Status:   domain.NewStatus("In Progress"),
			},
			want: BoardFetchUsecaseOutput{
				Boardid:  domain.NewId(1),
				Userid:   domain.NewUserid(1),
				Tabid:    domain.NewId(2),
				Priority: domain.NewPriority(1),
				Status:   domain.NewStatus("In Progress"),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := convert(tc.output)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
