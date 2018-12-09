package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func countDupes(s string) (hasTwo bool, hasThree bool) {
	letters := make(map[rune]int)

	for _, char := range s {
		letters[char]++
	}

	for _, count := range letters {
		if count == 2 {
			hasTwo = true
		}

		if count == 3 {
			hasThree = true
		}

		if hasTwo && hasThree {
			return
		}
	}

	return
}

func readAndCalcChecksum(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	twos := 0
	threes := 0

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		hasTwo, hasThree := countDupes(txt)

		if hasTwo {
			twos += 1
		}

		if hasThree {
			threes += 1
		}
	}

	return twos * threes
}

func main() {
	fmt.Printf("%v", readAndCalcChecksum(os.Stdin))
}
