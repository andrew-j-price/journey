package main

import "testing"

func TestAdd(t *testing.T) {

	got := add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
