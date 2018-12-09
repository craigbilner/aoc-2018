package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readAndTotal(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	total := 0

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		value, err := strconv.Atoi(txt)

		if err != nil {
			continue
		}

		total += value
	}

	return total
}

func main() {
	fmt.Printf("%v", readAndTotal(os.Stdin))
}
