package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPlay(t *testing.T) {
	input := `/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/  
`
	r := strings.NewReader(input)
	got := buildAndPlay(bufio.NewReader(r))

	if got != "7,3" {
		t.Errorf("Expected 7,3, got %v\n", got)
	}
}
