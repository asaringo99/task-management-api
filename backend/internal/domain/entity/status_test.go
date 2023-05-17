package domain

import (
	"reflect"
	"testing"
)

var (
	defaultStatus = Status{}
)

func TestNewStatus(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  Status
	}{
		{name: "CanCreateStatusType", value: "Pending", want: defaultStatus},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Status := NewStatus(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Status), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestStatusToValue(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "no diff", value: "In Progress", want: "In Progress"},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Status := NewStatus(tc.value)
		got := Status.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %s, got: %s", tc.name, tc.want, got)
		}
	}
}
