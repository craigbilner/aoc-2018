package main

import (
	"testing"
)

func TestReadAndSumMeta66(t *testing.T) {
	input := `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`
	result := rootValue(input)

	if result != 66 {
		t.Error("Expected 66 got ", result)
	}
}

func TestReadAndSumMeta30(t *testing.T) {
	input := `3 3 0 1 10 1 1 0 1 10 1 0 1 10 1 2 3`
	result := rootValue(input)

	if result != 30 {
		t.Error("Expected 30 got ", result)
	}
}
