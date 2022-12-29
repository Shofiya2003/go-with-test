package iterations

import "testing"

func TestIteration(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
