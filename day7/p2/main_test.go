package main

import (
	"strings"
	"testing"
)

func TestOrderedSteps(t *testing.T) {
	input := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`
	result := readAndOTimeSteps(0, 2, strings.NewReader(input))

	if result != 15 {
		t.Error("Expected 15 got ", result)
	}
}
