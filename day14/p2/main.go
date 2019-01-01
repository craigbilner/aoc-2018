package main

import (
	"fmt"
	"strconv"
	"strings"
)

type elf struct {
	indx int
}

func (e *elf) score(sb []int) int {
	return sb[e.indx]
}

func (e *elf) move(sb []int) {
	moves := 1 + sb[e.indx]

	e.indx = (e.indx + moves) % len(sb)
}

func newElf(indx int) *elf {
	return &elf{indx}
}

func howMany(scoreboard []int, after string) int {
	recipes := 0
	matching := false
	matchIndx := 0
	elf1, elf2 := newElf(0), newElf(1)

	for {
		if matchIndx == len(after) {
			recipes = len(scoreboard[:len(scoreboard)-len(after)])
			break
		}

		newRecipe := elf1.score(scoreboard) + elf2.score(scoreboard)
		newRecipes := strings.Split(fmt.Sprintf("%v", newRecipe), "")

		for _, nr := range newRecipes {
			v, _ := strconv.Atoi(nr)

			scoreboard = append(scoreboard, v)

			if !matching && nr == string(after[matchIndx]) {
				matching = true
				matchIndx++

				if matchIndx == len(after) {
					break
				}

				continue
			}

			if matching && nr == string(after[matchIndx]) {
				matchIndx++

				if matchIndx == len(after) {
					break
				}

				continue
			}

			if matching {
				matching = false
				matchIndx = 0
			}

			if !matching && nr == string(after[matchIndx]) {
				matching = true
				matchIndx++
			}
		}

		elf1.move(scoreboard)
		elf2.move(scoreboard)
	}

	return recipes
}

func main() {
	fmt.Printf("%v\n", howMany([]int{3, 7}, "084601"))
}
