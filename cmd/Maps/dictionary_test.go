//
package main

import (
	"testing"
)

//
func (t *testing.T) {
	dictionary := map[string]string{"test":"this is just a test"}

	got, err := Search(dictionary, "test")
	want := "this is just a test"

	if err != nil {
		t.errorf("an error occured")
	}

	if got != want {
		t.errorf("got %q want %q given, %q", got, want, test)
	}
}
