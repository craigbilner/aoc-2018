package main

import (
	"strings"
	"testing"
)

func TestReadAndCalcChecksum12(t *testing.T) {
	checksum := readAndCalcChecksum(strings.NewReader("abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab"))

	if checksum != 12 {
		t.Error("Expected 12 got ", checksum)
	}
}
