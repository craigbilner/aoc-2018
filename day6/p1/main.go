package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

type locationByDistances []*locationByDistance

func (lbd locationByDistances) Len() int {
	return len(lbd)
}

func (lbd locationByDistances) Less(i, j int) bool {
	return lbd[i].distance < lbd[j].distance
}

func (lbd locationByDistances) Swap(i, j int) {
	lbd[i], lbd[j] = lbd[j], lbd[i]
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

func findLargest(places []*coord) int {
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
	nearest := make(chan *locationByDistance, bx*by)

	for x := 0; x < bx; x++ {
		for y := 0; y < by; y++ {
			go func(tx, ty int) {
				lbds := make(chan *locationByDistance, len(places))
				target := &coord{tx, ty}

				go howFar(lbds, places, target)

				var xs locationByDistances
				for lbd := range lbds {
					xs = append(xs, lbd)
				}

				sort.Sort(xs)

				if xs[0].distance != xs[1].distance {
					nearest <- xs[0]
				}

				wg.Done()
			}(x, y)
		}
	}

	wg.Wait()
	close(nearest)

	locations := make(map[string][]*coord)
	for n := range nearest {
		id := fmt.Sprintf("%v%v", n.place.x, n.place.y)

		v, ok := locations[id]

		if ok && v == nil {
			continue
		}

		if n.target.x == 0 || n.target.y == 0 || n.target.x == bx || n.target.y == by {
			locations[id] = nil
			continue
		}

		if !ok {
			locations[id] = []*coord{n.target}
			continue
		}

		locations[id] = append(v, n.target)
	}

	la := 0
	for _, v := range locations {
		if len(v) > la {
			la = len(v)
		}
	}

	return la
}

func largestArea(r io.Reader) int {
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

	return findLargest(places)
}

func main() {
	fmt.Printf("%v", largestArea(os.Stdin))
}
