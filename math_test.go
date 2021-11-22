package main

import "testing"

func testAdd(t *testing.T) {

	got := add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
