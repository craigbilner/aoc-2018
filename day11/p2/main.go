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
	x, y, size, power int
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
	ch := make(chan *cell)
	var wg sync.WaitGroup

	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			wg.Add(1)
			go func(x, y int) {
				s := 1
				for {
					if x+s > 301 || y+s > 301 {
						break
					}

					total := 0
					for k := 0; k < s; k++ {
						for l := 0; l < s; l++ {
							v, _ := g[fmt.Sprintf("%v,%v", x+l, y+k)]
							total += v
						}
					}

					ch <- &cell{x, y, s, total}
					s++
				}
				wg.Done()
			}(j, i)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	highest := &cell{}
	count := 0
	for c := range ch {
		count++

		if count%1000 == 0 {
			fmt.Printf("%v\n", count)
		}

		if c.power > highest.power {
			highest = c
		}
	}

	return fmt.Sprintf("%v,%v,%v with %v", highest.x, highest.y, highest.size, highest.power)
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
