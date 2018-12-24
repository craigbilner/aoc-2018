package main

import "testing"

func TestLargestPower18(t *testing.T) {
	got := largestPower(18)

	if got != "33,45 with 29" {
		t.Errorf("Expected 33,45 with 29, got %v\n", got)
	}
}

func TestLargestPower42(t *testing.T) {
	got := largestPower(42)

	if got != "21,61 with 30" {
		t.Errorf("Expected 21,61 with 30, got %v\n", got)
	}
}
