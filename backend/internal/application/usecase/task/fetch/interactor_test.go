package usecase

import (
	"reflect"
	"testing"

	repository "github.com/asaringo99/task_management/internal/application/repository/task/fetch"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name   string
		output repository.TaskFetchRepositoryOutput
		want   TaskFetchUsecaseOutput
	}{
		{
			name: "no diff1",
			output: repository.TaskFetchRepositoryOutput{
				Taskid:   domain.NewTaskid(1),
				Userid:   domain.NewUserid(1),
				Status:   domain.NewStatus("Pending"),
				Contents: domain.NewContents("test"),
				Priority: domain.NewPriority(1),
			},
			want: TaskFetchUsecaseOutput{
				Taskid:   domain.NewTaskid(1),
				Userid:   domain.NewUserid(1),
				Status:   domain.NewStatus("Pending"),
				Contents: domain.NewContents("test"),
				Priority: domain.NewPriority(1),
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
