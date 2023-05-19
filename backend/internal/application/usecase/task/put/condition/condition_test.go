package condition

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/entity"
)

func TestNewTaskPutCondition(t *testing.T) {
	tests := []struct {
		name     string
		taskid   int
		userid   int
		status   string
		priority int
		contents string
		want     TaskPutCondition
	}{
		{
			name:     "no diff1",
			taskid:   1,
			userid:   1,
			status:   "Pending",
			contents: "test",
			priority: 1,
			want: TaskPutCondition{
				taskid:   domain.NewTaskid(1),
				userid:   domain.NewUserid(1),
				status:   domain.NewStatus("Pending"),
				contents: domain.NewContents("test"),
				priority: domain.NewPriority(1),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewTaskPutCondition(tc.taskid, tc.userid, tc.status, tc.priority, tc.contents)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
