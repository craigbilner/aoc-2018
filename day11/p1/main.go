package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type cell struct {
	x, y, power int
}

func getHundred(i int) int {
	return int(math.Floor(float64(i)/100.) - math.Floor(float64(i)/1000.)*10.)
}

func powerLevel(x, y, serialNumber int) int {
	rackId := x + 10
	return getHundred((rackId*y+serialNumber)*rackId) - 5
}

func newGrid(serialNumber int) map[string]int {
	m := make(map[string]int)
	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			m[fmt.Sprintf("%v,%v", j, i)] = powerLevel(j, i, serialNumber)
		}
	}
	return m
}

func largestPower(serialNumber int) string {
	g := newGrid(serialNumber)
	ch := make(chan *cell, 88804)
	var wg sync.WaitGroup
	wg.Add(88804)

	for i := 1; i < 299; i++ {
		for j := 1; j < 299; j++ {
			go func(x, y int) {
				total := 0
				for k := 0; k < 3; k++ {
					for l := 0; l < 3; l++ {
						v, _ := g[fmt.Sprintf("%v,%v", x+l, y+k)]
						total += v
					}
				}

				ch <- &cell{x, y, total}
				wg.Done()
			}(j, i)
		}
	}

	wg.Wait()
	close(ch)

	highest := &cell{}
	for c := range ch {
		if c.power > highest.power {
			highest = c
		}
	}

	return fmt.Sprintf("%v,%v with %v", highest.x, highest.y, highest.power)
}

func unsafeAtoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	fmt.Printf(largestPower(unsafeAtoi(strings.Trim(text, "\n"))) + "\n")
}
