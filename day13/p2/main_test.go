package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func TestPlay(t *testing.T) {
	input := `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
`

	r := strings.NewReader(input)
	got := buildAndPlay(bufio.NewReader(r))

	if got != "6,4" {
		t.Errorf("Expected 6,4, got %v\n", got)
	}
}

func TestSort(t *testing.T) {
	carts := []*cart{
		newCart(left, &rail{what: 0, coord: &coord{3, 4}}),
		newCart(left, &rail{what: 0, coord: &coord{5, 6}}),
		newCart(left, &rail{what: 0, coord: &coord{2, 6}}),
		newCart(left, &rail{what: 0, coord: &coord{2, 3}}),
	}

	sort.Sort(cartByCoord(carts))

	if carts[0].x != 2 && carts[0].y != 3 {
		t.Errorf("Expected 2,3 got %v,%v", carts[0].x, carts[0].y)
	}

	if carts[2].x != 2 && carts[2].y != 6 {
		t.Errorf("Expected 2,6 got %v,%v", carts[2].x, carts[2].y)
	}
}
