package main

import (
	"strings"
	"testing"
)

func TestReadAndFindNoOverlap3(t *testing.T) {
	id := readAndFindNoOverlap(strings.NewReader("#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"))

	if id != 3 {
		t.Error("Expected 3 got ", id)
	}
}
