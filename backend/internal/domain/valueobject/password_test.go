package domain

import (
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var (
	defaultPassword = Password{}
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  Password
	}{
		{name: "CanCreatePasswordType", value: "test", want: defaultPassword},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Password := NewPassword(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Password), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestPasswordToValue(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "no diff", value: "test", want: "test"},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Password := NewPassword(tc.value)
		got := Password.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %s, got: %s", tc.name, tc.want, got)
		}
	}
}

func TestNewHashedPassword(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  Password
	}{
		{name: "CanCreatePasswordType", value: "test", want: defaultPassword},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Password := NewHashedPassword(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Password), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestHashedPasswordToValue(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "no diff", value: "test", want: "test"},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Password := NewHashedPassword(tc.value)
		got := Password.ToValue()
		if err := bcrypt.CompareHashAndPassword([]byte(got), []byte(tc.want)); err != nil {
			t.Fatalf("Hash Fault")
		}
	}
}
