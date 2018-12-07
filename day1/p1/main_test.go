package main

import (
	"strings"
	"testing"
)

func TestReadAndTotal(t *testing.T) {
	total := readAndTotal(strings.NewReader("+13\n+12\n-4"))

	if total != 21 {
		t.Error("Expected 21 got ", total)
	}
}
