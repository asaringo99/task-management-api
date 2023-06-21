package presenter

import (
	"reflect"
	"testing"

	usecase "github.com/asaringo99/task_management/internal/application/usecase/board/create"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name   string
		output usecase.BoardCreateUsecaseOutput
		want   BoardCreatePresenterOutputDto
	}{
		{
			name: "no diff1",
			output: usecase.BoardCreateUsecaseOutput{
				Boardid:  domain.NewId(1),
				Userid:   domain.NewUserid(1),
				Tabid:    domain.NewId(2),
				Priority: domain.NewPriority(1),
				Status:   domain.NewStatus("test"),
			},
			want: BoardCreatePresenterOutputDto{
				Boardid:  1,
				Userid:   1,
				Tabid:    2,
				Priority: 1,
				Status:   "test",
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
