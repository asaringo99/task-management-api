package presenter

import (
	"reflect"
	"testing"

	usecase "github.com/asaringo99/task_management/internal/application/usecase/task/fetch"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name   string
		output usecase.TaskFetchUsecaseOutput
		want   TaskFetchPresenterOutputDto
	}{
		{
			name: "no diff1",
			output: usecase.TaskFetchUsecaseOutput{
				Taskid:   domain.NewTaskid(1),
				Userid:   domain.NewUserid(1),
				Status:   domain.NewStatus("Pending"),
				Priority: domain.NewPriority(1),
				Contents: domain.NewContents("test"),
			},
			want: TaskFetchPresenterOutputDto{
				Taskid:   1,
				Userid:   1,
				Status:   "Pending",
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
