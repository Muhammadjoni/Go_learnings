package main

import "testing"


func TestMain(t *testing.T) {

	got := main.
	want := "Don't communicate by sharing memory, share memory by communicating."

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
