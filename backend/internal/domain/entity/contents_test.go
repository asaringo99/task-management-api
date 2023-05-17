package domain

import (
	"reflect"
	"testing"
)

var (
	defaultContents = Contents{}
)

func TestNewContents(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  Contents
	}{
		{name: "CanCreateContentsType", value: "test", want: defaultContents},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		contents := NewContents(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(contents), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestContentsToValue(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "no diff", value: "test", want: "test"},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		contents := NewContents(tc.value)
		got := contents.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %s, got: %s", tc.name, tc.want, got)
		}
	}
}
