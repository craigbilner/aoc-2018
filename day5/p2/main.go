package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

func reacts(charCodeA, charCodeB uint8) bool {
	return math.Abs(float64(int(charCodeA)-int(charCodeB))) == 32
}

func reduce(polymer string) string {
	if len(polymer) <= 1 {
		return polymer
	}

	if len(polymer) == 2 {
		if reacts(polymer[0], polymer[1]) {
			return ""
		}

		return polymer
	}

	ru := polymer
	i := 0

	for {
		if i >= len(ru)-1 {
			break
		}

		if !reacts(ru[i], ru[i+1]) {
			i++
			continue
		}

		if i == 0 {
			ru = ru[2:]
			i = 0
			continue
		}

		ru = ru[:i] + ru[i+2:]
		i = i - 1
	}

	return ru
}

func remainingUnits(polymer string) string {
	seen := make(map[int32]struct{})
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, r := range strings.ToLower(polymer) {
		if _, ok := seen[r]; ok {
			continue
		}

		seen[r] = struct{}{}

		wg.Add(1)
		go func(char rune) {
			re := regexp.MustCompile(fmt.Sprintf("(?i)%c", char))

			ch <- reduce(re.ReplaceAllString(polymer, ""))
			wg.Done()
		}(rune(r))
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	best := polymer
	for v := range ch {
		if len(v) < len(best) {
			best = v
		}
	}

	return best
}

func main() {
	pwd, _ := os.Getwd()
	txt, _ := ioutil.ReadFile(filepath.Join(pwd, "input.txt"))

	println(len(remainingUnits(string(txt))))
}
