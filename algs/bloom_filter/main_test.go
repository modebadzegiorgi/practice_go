package main

import "testing"

func TestHash(t *testing.T) {
	got := hash("Piano")
	want := 684

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
