package main

import "testing"

func TestRemainingUnits4(t *testing.T) {
	got := remainingUnits("dabAcCaCBAcCcaDA")

	if got != "daDA" {
		t.Errorf("Wanted daDA, got %v", got)
	}
}
