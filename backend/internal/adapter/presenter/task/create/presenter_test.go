package presenter

import (
	"reflect"
	"testing"

	usecase "github.com/asaringo99/task_management/internal/application/usecase/task/create"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name   string
		output usecase.TaskCreateUsecaseOutput
		want   TaskCreatePresenterOutputDto
	}{
		{
			name: "no diff1",
			output: usecase.TaskCreateUsecaseOutput{
				Taskid:   domain.NewTaskid(1),
				Userid:   domain.NewUserid(1),
				Boardid:  domain.NewId(2),
				Priority: domain.NewPriority(1),
				Contents: domain.NewContents("test"),
			},
			want: TaskCreatePresenterOutputDto{
				Taskid:   1,
				Userid:   1,
				Boardid:  2,
				Priority: 1,
				Contents: "test",
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
