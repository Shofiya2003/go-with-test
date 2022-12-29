package main

import "testing"

func TestGreeting(t *testing.T) {
	t.Run("say good morning", func(t *testing.T) {
		got := Greet("chris", 10)
		want := "good morning chris"
		assertMessage(t, got, want)
	})

	t.Run("say good afternoon", func(t *testing.T) {
		got := Greet("chris", 13)
		want := "good afternoon chris"
		assertMessage(t, got, want)
	})

	t.Run("say good evening", func(t *testing.T) {
		got := Greet("chris", 17)
		want := "good evening chris"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
