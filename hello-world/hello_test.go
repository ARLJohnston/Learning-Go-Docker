package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("Saying hello to myself", func(t *testing.T) {
    got := hello("Alistair")
    want := "Hello, Alistair!"

    if got != want {
            t.Errorf("got %q want %q", got, want)
    }
  })


  t.Run("Saying hello world when no one answers", func(t *testing.T) {
    got := hello("")
    want := "Hello, World!"

    if got != want {
            t.Errorf("got %q want %q", got, want)
    }
  })
}
