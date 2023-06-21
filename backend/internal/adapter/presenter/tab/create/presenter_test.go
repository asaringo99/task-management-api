package presenter

import (
	"reflect"
	"testing"

	usecase "github.com/asaringo99/task_management/internal/application/usecase/tab/create"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name   string
		output usecase.TabCreateUsecaseOutput
		want   TabCreatePresenterOutputDto
	}{
		{
			name: "no diff1",
			output: usecase.TabCreateUsecaseOutput{
				Tabid:    domain.NewId(1),
				Userid:   domain.NewUserid(1),
				IsActive: true,
				Title:    domain.NewTitle("test"),
			},
			want: TabCreatePresenterOutputDto{
				Tabid:    1,
				Userid:   1,
				IsActive: true,
				Title:    "test",
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
