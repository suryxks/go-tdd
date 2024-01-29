package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	CountDown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
