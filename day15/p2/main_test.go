package main

import (
	"strings"
	"testing"
)

func TestReadAndFight4v2(t *testing.T) {
	input := `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
`
	r := strings.NewReader(input)
	want := 4988
	got := fight(readerToArena(r))

	if got != want {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

func TestReadAndFight3v6(t *testing.T) {
	input := `#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######
`
	r := strings.NewReader(input)
	want := 31284
	got := fight(readerToArena(r))

	if got != want {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

func TestReadAndFight5v2(t *testing.T) {
	input := `#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######
`
	r := strings.NewReader(input)
	want := 3478
	got := fight(readerToArena(r))

	if got != want {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

func TestReadAndFight4v2b(t *testing.T) {
	input := `#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######
`
	r := strings.NewReader(input)
	want := 6474
	got := fight(readerToArena(r))

	if got != want {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

func TestReadAndFight5v1(t *testing.T) {
	input := `#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########
`
	r := strings.NewReader(input)
	want := 1140
	got := fight(readerToArena(r))

	if got != want {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}
