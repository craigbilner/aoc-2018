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

func score(scoreboard []int, made int) string {
	s := ""
	elf1, elf2 := newElf(0), newElf(1)

	for {
		if len(scoreboard) >= (made + 10) {
			for _, c := range scoreboard[made : made+10] {
				s += strconv.Itoa(c)
			}

			break
		}

		newRecipe := elf1.score(scoreboard) + elf2.score(scoreboard)
		newRecipes := strings.Split(fmt.Sprintf("%v", newRecipe), "")

		for _, nr := range newRecipes {
			v, _ := strconv.Atoi(nr)
			scoreboard = append(scoreboard, v)
		}

		elf1.move(scoreboard)
		elf2.move(scoreboard)
	}

	return s
}

func main() {
	fmt.Printf("%v\n", score([]int{3, 7}, 84601))
}
