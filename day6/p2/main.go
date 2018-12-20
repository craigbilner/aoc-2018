package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type coord struct {
	x, y int
}

type locationByDistance struct {
	target   *coord
	place    *coord
	distance int
}

func strToCoord(s string) *coord {
	x := strings.Split(s, ", ")
	x1, _ := strconv.Atoi(x[0])
	x2, _ := strconv.Atoi(x[1])

	return &coord{
		x1, x2,
	}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func distance(c1, c2 *coord) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

func howFar(out chan<- *locationByDistance, places []*coord, target *coord) {
	for _, place := range places {
		out <- &locationByDistance{target, place, distance(place, target)}
	}

	close(out)
}

func findLargest(limit int, places []*coord) int {
	var bx, by int
	for _, c := range places {
		if c.x > bx {
			bx = c.x
		}

		if c.y > by {
			by = c.y
		}
	}
	bx++
	by++

	var wg sync.WaitGroup
	wg.Add(bx * by)
	withinLimit := make(chan *coord, bx*by)

	for x := 0; x < bx; x++ {
		for y := 0; y < by; y++ {
			go func(tx, ty int) {
				lbds := make(chan *locationByDistance, len(places))
				target := &coord{tx, ty}

				go howFar(lbds, places, target)

				sum := 0
				for lbd := range lbds {
					sum += lbd.distance
				}

				if sum < limit {
					withinLimit <- target
				}

				wg.Done()
			}(x, y)
		}
	}

	wg.Wait()
	close(withinLimit)

	count := 0
	for range withinLimit {
		count++
	}

	return count
}

func largestArea(limit int, r io.Reader) int {
	scanner := bufio.NewScanner(r)
	var places []*coord

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		coord := strToCoord(txt)
		places = append(places, coord)
	}

	return findLargest(limit, places)
}

func main() {
	fmt.Printf("%v", largestArea(10000, os.Stdin))
}
