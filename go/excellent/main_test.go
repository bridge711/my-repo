package main

import "testing"

func TestEven0r0dd(t *testing.T) {
	result := Even0r0dd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}
