package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alistair")

	got := buffer.String()
	want := "Hello, Alistair"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
