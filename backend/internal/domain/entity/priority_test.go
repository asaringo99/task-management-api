package domain

import (
	"reflect"
	"testing"
)

var (
	defaultPriority = Priority{}
)

func TestNewPriority(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  Priority
	}{
		{name: "CanCreatePriorityType", value: 1, want: defaultPriority},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Priority := NewPriority(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Priority), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestPriorityToValue(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "no diff", value: 1, want: 1},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Priority := NewPriority(tc.value)
		got := Priority.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %d, got: %d", tc.name, tc.want, got)
		}
	}
}
