package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to a person", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello Chris"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say `Hello World` when name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say in spanish", func(t *testing.T) {
		got := Hello("Chris", "spanish")
		want := "Hola Chris"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say in french", func(t *testing.T) {
		got := Hello("Michel", "french")
		want := "Bonjour Michel"
		asserCorrectMessage(t, got, want)
	})

}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")
	got := buffer.String()
	want := "Hello Chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func asserCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
