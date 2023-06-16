package domain

import (
	"reflect"
	"testing"
)

var (
	defaultUserid   = Userid{}
	defaultUsername = Username{}
)

func TestNewUserid(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  Userid
	}{
		{name: "CanCreateUseridType", value: 1, want: defaultUserid},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Userid := NewUserid(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Userid), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestUseridToValue(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "no diff", value: 3, want: 3},
		{name: "no diff", value: 25, want: 25},
		{name: "no diff", value: 100, want: 100},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Userid := NewUserid(tc.value)
		got := Userid.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %d, got: %d", tc.name, tc.want, got)
		}
	}
}

func TestNewUsername(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  Username
	}{
		{name: "CanCreateUsernameType", value: "user1", want: defaultUsername},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Username := NewUsername(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Username), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestUsernameToValue(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "no diff", value: "user1", want: "user1"},
		{name: "no diff", value: "user10", want: "user10"},
		{name: "no diff", value: "test", want: "test"},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Username := NewUsername(tc.value)
		got := Username.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %s, got: %s", tc.name, tc.want, got)
		}
	}
}
