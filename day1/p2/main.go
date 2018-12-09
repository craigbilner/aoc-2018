package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func findDupe(seen map[int]struct{}, total int, values *[]int) int {
	for _, value := range *values {
		total += value

		if _, ok := seen[total]; ok {
			return total
		}

		seen[total] = struct{}{}
	}

	return findDupe(seen, total, values)
}

func readAndFindDupe(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	seen := make(map[int]struct{})
	total := 0
	var values []int

	seen[0] = struct{}{}

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

		if _, ok := seen[total]; ok {
			return total
		}

		seen[total] = struct{}{}
		values = append(values, value)
	}

	return findDupe(seen, total, &values)
}

func main() {
	fmt.Printf("%v", readAndFindDupe(os.Stdin))
}
