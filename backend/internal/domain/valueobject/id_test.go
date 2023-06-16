package domain

import (
	"reflect"
	"testing"
)

var (
	defaultId = Id{}
)

func TestNewId(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  Id
	}{
		{name: "CanCreateIdType", value: 1, want: defaultId},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Id := NewId(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Id), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestIdToValue(t *testing.T) {
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
		Id := NewId(tc.value)
		got := Id.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %d, got: %d", tc.name, tc.want, got)
		}
	}
}
