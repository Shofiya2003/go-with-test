package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

var write string = "write"
var sleep string = "sleep"

func TestCountdown(t *testing.T) {

	t.Run("sleep 3 times", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q ", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enought call to the sleeper want 3 got %d ", spySleeper.Calls)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperation{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted %v got %v ", want, spySleepPrinter.Calls)
		}

	})
}
