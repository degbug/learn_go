package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("degbug")
	want := "Hello, world1"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
