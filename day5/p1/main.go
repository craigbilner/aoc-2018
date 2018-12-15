package main

import (
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
)

func reacts(charCodeA, charCodeB uint8) bool {
	return math.Abs(float64(int(charCodeA)-int(charCodeB))) == 32
}

func remainingUnits(polymer string) string {
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

func main() {
	pwd, _ := os.Getwd()
	txt, _ := ioutil.ReadFile(filepath.Join(pwd, "input.txt"))

	println(len(remainingUnits(string(txt))))
}
