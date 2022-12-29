package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to a person", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say `Hello World` when name is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		asserCorrectMessage(t, got, want)
	})

}

func asserCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
