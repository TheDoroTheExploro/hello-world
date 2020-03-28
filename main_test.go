package main

import "testing"

func TestMultiply(t *testing.T) {
	res := Multiply(3)
	if res != 9 {
		t.Error()
	}
}
