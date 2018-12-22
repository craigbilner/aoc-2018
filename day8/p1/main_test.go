package main

import (
	"testing"
)

func TestReadAndSumMeta138(t *testing.T) {
	input := `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`
	result := sumMeta(input)

	if result != 138 {
		t.Error("Expected 138 got ", result)
	}
}

func TestReadAndSumMeta50(t *testing.T) {
	input := `3 1 0 1 10 1 1 0 1 10 10 0 1 10 10`
	result := sumMeta(input)

	if result != 50 {
		t.Error("Expected 50 got ", result)
	}
}
