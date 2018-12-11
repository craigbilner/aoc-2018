package main

import (
	"strings"
	"testing"
)

func TestReadAndCountOverlap4(t *testing.T) {
	count := readAndCountOverlap(strings.NewReader("#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"))

	if count != 4 {
		t.Error("Expected 4 got ", count)
	}
}
