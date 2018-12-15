package main

import "testing"

func TestRemainingUnits0(t *testing.T) {
	got := remainingUnits("aA")

	if got != "" {
		t.Errorf("Wanted nothing, got %v", got)
	}
}

func TestRemainingUnits00(t *testing.T) {
	got := remainingUnits("abBA")

	if got != "" {
		t.Errorf("Wanted nothing, got %v", got)
	}
}

func TestRemainingUnits4(t *testing.T) {
	got := remainingUnits("abAB")

	if got != "abAB" {
		t.Errorf("Wanted abAB, got %v", got)
	}
}

func TestRemainingUnits6(t *testing.T) {
	got := remainingUnits("aabAAB")

	if got != "aabAAB" {
		t.Errorf("Wanted aabAAB, got %v", got)
	}
}

func TestRemainingUnits10(t *testing.T) {
	got := remainingUnits("dabAcCaCBAcCcaDA")

	if got != "dabCBAcaDA" {
		t.Errorf("Wanted dabCBAcaDA, got %v", got)
	}
}

func TestRemainingUnitsTriples(t *testing.T) {
	got := remainingUnits("abcCBA")

	if got != "" {
		t.Errorf("Wanted nothing, got %v", got)
	}
}

func TestRemainingUnitsDoubles(t *testing.T) {
	got := remainingUnits("abcbdcCDBefg")

	if got != "abcefg" {
		t.Errorf("Wanted abcefg, got %v", got)
	}
}
