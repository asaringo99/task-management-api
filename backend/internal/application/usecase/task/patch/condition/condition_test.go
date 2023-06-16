package condition

import (
	"reflect"
	"testing"

	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestNewTaskPutCondition(t *testing.T) {
	tests := []struct {
		name   string
		taskid int
		userid int
		fix    PatchData
		want   TaskPatchCondition
	}{
		{
			name:   "no diff1",
			taskid: 1,
			userid: 1,
			fix: PatchData{
				Column: "boardid",
				Value:  2,
			},
			want: TaskPatchCondition{
				taskid: domain.NewTaskid(1),
				userid: domain.NewUserid(1),
				fix: PatchData{
					Column: "boardid",
					Value:  2,
				},
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewTaskPatchCondition(tc.taskid, tc.userid, tc.fix)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
