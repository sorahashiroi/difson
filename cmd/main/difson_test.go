package main

import "testing"

func Example_difson() {
	goMain([]string{"difson"})
	// Output:
	// Welcome to Difson!
}

func TestHello(t *testing.T) {
	got := hello()
	want := "Welcome to Difson!"
	if got != want {
			t.Errorf("hello() = %q, want %q", got, want)
	}
}