package presenter

import (
	"reflect"
	"testing"

	usecase "github.com/asaringo99/task_management/internal/application/usecase/board/fetch"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name   string
		output usecase.BoardFetchUsecaseOutput
		want   BoardFetchPresenterOutputDto
	}{
		{
			name: "no diff1",
			output: usecase.BoardFetchUsecaseOutput{
				Boardid:  domain.NewId(1),
				Userid:   domain.NewUserid(1),
				Priority: domain.NewPriority(1),
				Status:   domain.NewStatus("test"),
			},
			want: BoardFetchPresenterOutputDto{
				Boardid:  1,
				Userid:   1,
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
