package main

import (
	"strings"
	"testing"
)

func TestReadAndFindDupe0(t *testing.T) {
	dupe := readAndFindDupe(strings.NewReader("+1\n-1"))

	if dupe != 0 {
		t.Error("Expected 0 got ", dupe)
	}
}

func TestReadAndFindDupe10(t *testing.T) {
	dupe := readAndFindDupe(strings.NewReader("+3\n+3\n+4\n-2\n-4"))

	if dupe != 10 {
		t.Error("Expected 10 got ", dupe)
	}
}

func TestReadAndFindDupe5(t *testing.T) {
	dupe := readAndFindDupe(strings.NewReader("-6\n+3\n+8\n+5\n-6"))

	if dupe != 5 {
		t.Error("Expected 5 got ", dupe)
	}
}

func TestReadAndFindDupe14(t *testing.T) {
	dupe := readAndFindDupe(strings.NewReader("+7\n+7\n-2\n-7\n-4"))

	if dupe != 14 {
		t.Error("Expected 14 got ", dupe)
	}
}
