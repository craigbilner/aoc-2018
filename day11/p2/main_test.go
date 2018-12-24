package main

import "testing"

func TestLargestPower18(t *testing.T) {
	got := largestPower(18)

	if got != "90,269,16 with 113" {
		t.Errorf("Expected 90,269,16 with 113, got %v\n", got)
	}
}

func TestLargestPower42(t *testing.T) {
	got := largestPower(42)

	if got != "232,251,12 with 119" {
		t.Errorf("Expected 232,251,12 with 119, got %v\n", got)
	}
}
