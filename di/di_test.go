package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Lorenas")

	got := buffer.String()
	want := "Hello, Lorenas"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}