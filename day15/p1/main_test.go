package main

import (
	"strings"
	"testing"
)

func TestFirstSteps1(t *testing.T) {
	input := `#######
#E..G.#
#...#.#
#.G.#G#
#######
`
	r := strings.NewReader(input)
	a, elves, goblins := readerToArena(r)

	positions := make(map[string]*unit)
	for _, v := range elves {
		positions[v.coord.String()] = v
	}
	for _, v := range goblins {
		positions[v.coord.String()] = v
	}

	sd, fs, ok := firstSteps(1, a, positions, &coord{1, 1}, &coord{1, 3})

	if ok != true {
		t.Errorf("Expected success")
	}

	if sd != 2 {
		t.Errorf("Expected a distance of %v, got %v", 2, sd)
	}

	if len(fs) != 1 {
		t.Fatalf("Wrong number of routes, expected %v, got %v", 1, len(fs))
	}

	if fs[0].x != 1 && fs[0].y != 2 {
		t.Errorf("First step should be %v,%v, got %v,%v", 1, 2, fs[0].x, fs[0].y)
	}
}

func TestFirstSteps3(t *testing.T) {
	input := `#######
#E..G.#
#...#.#
#.G.#G#
#######
`
	r := strings.NewReader(input)
	a, elves, goblins := readerToArena(r)

	positions := make(map[string]*unit)
	for _, v := range elves {
		positions[v.coord.String()] = v
	}
	for _, v := range goblins {
		positions[v.coord.String()] = v
	}

	sd, fs, ok := firstSteps(1, a, positions, &coord{1, 1}, &coord{3, 3})

	if ok != true {
		t.Errorf("Expected success")
	}

	if sd != 4 {
		t.Errorf("Expected a distance of %v, got %v", 4, sd)
	}

	if len(fs) != 2 {
		t.Fatalf("Wrong number of routes, expected %v, got %v", 2, len(fs))
	}

	if fs[0].x != 2 && fs[0].y != 1 {
		t.Errorf("First step should be %v,%v, got %v,%v", 2, 1, fs[0].x, fs[0].y)
	}

	if fs[1].x != 1 && fs[1].y != 2 {
		t.Errorf("Third step should be %v,%v, got %v,%v", 2, 1, fs[1].x, fs[1].y)
	}
}

func TestFirstStepsImpossible(t *testing.T) {
	input := `#######
#E..G.#
#...#.#
#.G.#G#
#######
`
	r := strings.NewReader(input)
	a, elves, goblins := readerToArena(r)

	positions := make(map[string]*unit)
	for _, v := range elves {
		positions[v.coord.String()] = v
	}
	for _, v := range goblins {
		positions[v.coord.String()] = v
	}

	sd, fs, ok := firstSteps(1, a, positions, &coord{1, 1}, &coord{5, 1})

	if sd != 0 {
		t.Errorf("Expected a distance of %v, got %v", 0, sd)
	}

	if len(fs) != 0 {
		t.Fatalf("Too many routes, expected %v, got %v", 0, len(fs))
	}

	if ok != false {
		t.Errorf("Expected failure")
	}
}

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
	want := 27730
	got := fight(readerToArena(r))

	if got != want {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

func TestReadAndFight2v6(t *testing.T) {
	input := `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######
`
	r := strings.NewReader(input)
	want := 36334
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
	want := 39514
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
	want := 27755
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
	want := 28944
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
	want := 18740
	got := fight(readerToArena(r))

	if got != want {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}
