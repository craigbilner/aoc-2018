package main

import (
	"strings"
	"testing"
)

func TestLargestArea17(t *testing.T) {
	input := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`
	result := largestArea(32, strings.NewReader(input))

	if result != 16 {
		t.Error("Expected 16 got ", result)
	}
}
