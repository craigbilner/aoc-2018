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

		operator := string(txt[0])
		value, _ := strconv.Atoi(txt[1:])

		switch operator {
		case "+":
			total += value
		case "-":
			total -= value
		default:
			continue
		}
	}

	return total
}

func main() {
	fmt.Printf("%v", readAndTotal(os.Stdin))
}
