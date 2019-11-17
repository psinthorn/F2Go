package main

import "testing"

func TestMain(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say Hello", func(t *testing.T) {
		got := Hello()
		want := "Hello F2Code"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Get http response", func(t *testing.T) {
		got := "200"
		want := "200"
		assertCorrectMessage(t, got, want)
	})

}
