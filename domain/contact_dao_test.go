package domain

import "testing"

func TestGetUserUserNotFound(t *testing.T) {
	got := GetUser(0)
	want := ""

	if got != want {
		t.Error("User not found")
	}
}
