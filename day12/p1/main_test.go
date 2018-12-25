package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadAndSumPlantPots(t *testing.T) {
	input := `initial state: ##.......#.######.##..#...#.#.#..#...#..####..#.##...#....#...##..#..#.##.##.###.##.#.......###....#

.#### => .
....# => .
###.. => .
..#.# => .
##### => #
####. => .
#.##. => #
#.#.# => .
##.#. => #
.###. => .
#..#. => #
###.# => .
#.### => .
##... => #
.#.## => .
..#.. => .
#...# => #
..... => .
.##.. => .
...#. => .
#.#.. => .
.#..# => #
.#.#. => .
.#... => #
..##. => .
#..## => .
##.## => #
...## => #
..### => #
#.... => .
.##.# => #
##..# => #
`
	r := strings.NewReader(input)
	got := readAndSumPlantPots(20, bufio.NewReader(r))

	if got != 2840 {
		t.Errorf("Expected 2840, got %v\n", got)
	}
}