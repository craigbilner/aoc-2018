package main

import "testing"

func TestInterpolate(t *testing.T) {
	got := interpolate(2000)

	if got != 81684 {
		t.Errorf("Expected 81684, got %v\n", got)
	}
}
