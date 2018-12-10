package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func compare(s1 string, s2 string) (matched string, ok bool) {
	diff := 0
	indx := 0

	for i := range s1 {
		if s1[i] != s2[i] {
			indx = i
			diff++
		}

		if diff > 1 {
			break
		}
	}

	if diff != 1 {
		return "", false
	}

	return s1[:indx] + s1[indx+1:], true
}

func findMatch(out chan<- string, words []string, test string) {
	for _, word := range words {
		matched, ok := compare(test, word)

		if !ok {
			continue
		}

		out <- matched
		break
	}
}

func readAndCalcCommonString(r io.Reader) string {
	matched := make(chan string, 1)
	scanner := bufio.NewScanner(r)
	var words []string

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		if len(words) == 0 {
			words = append(words, txt)
			continue
		}

		go findMatch(matched, words, txt)

		words = append(words, txt)
	}

	return <-matched
}

func main() {
	fmt.Printf("%s", readAndCalcCommonString(os.Stdin))
}
