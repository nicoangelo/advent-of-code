package main

import "testing"

func TestRead_file(t *testing.T) {
	res := read_file("./example")
	if res != 198 {
		t.Fatalf("Epsilon wrong, got: %d, want: %d", res, 198)
	}
}
