package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lorenas")
	want := "Hello, Lorenas"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
