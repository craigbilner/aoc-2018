package main

import (
	"strings"
	"testing"
)

func TestReadAndFindDupe0(t *testing.T) {
	common := readAndCalcCommonString(strings.NewReader("abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz"))

	if common != "fgij" {
		t.Error("Expected fgij got ", common)
	}
}
