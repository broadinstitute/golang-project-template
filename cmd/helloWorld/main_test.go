package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, World!"
	got := hello()
	if got != want {
		t.Errorf("Expected: %q, got: %q", want, got)
	}
}

func TestDoSomething(t *testing.T) {
	// adding a do nothing test for demonstration purposes
	doSomething()
}
