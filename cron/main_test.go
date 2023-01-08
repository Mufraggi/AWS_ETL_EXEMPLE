package main

import "testing"

func TestAbs(t *testing.T) {
	a := 1
	if a != 1 {
		t.Errorf("Abs(-1) = %d; want 1", a)
	}
}
