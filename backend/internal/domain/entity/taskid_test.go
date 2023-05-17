package domain

import (
	"reflect"
	"testing"
)

var (
	defaultTaskid = Taskid{}
)

func TestNewTaskid(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  Taskid
	}{
		{name: "CanCreateTaskidType", value: 1, want: defaultTaskid},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Taskid := NewTaskid(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Taskid), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestTaskidToValue(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "no diff", value: 1, want: 1},
		{name: "no diff", value: 10, want: 10},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Taskid := NewTaskid(tc.value)
		got := Taskid.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %d, got: %d", tc.name, tc.want, got)
		}
	}
}
