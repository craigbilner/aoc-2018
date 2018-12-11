package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

func noOverlap(out chan<- int, fabric [][]int, c claim) {
	hasNoOverlap := true
Loop:
	for i := c.top; i < c.height+c.top; i++ {
		for j := c.left; j < c.width+c.left; j++ {
			if fabric[i][j] > 1 {
				hasNoOverlap = false
				break Loop
			}
		}
	}

	if hasNoOverlap {
		out <- c.id
	}
}

func findNoOverlap(width int, height int, claims []claim) int {
	fabric := make([][]int, height)
	for i := range fabric {
		fabric[i] = make([]int, width)
	}

	for _, claim := range claims {
		for i := claim.top; i < claim.height+claim.top; i++ {
			for j := claim.left; j < claim.width+claim.left; j++ {
				fabric[i][j]++
			}
		}
	}

	id := make(chan int, 1)

	for _, v := range claims {
		go func(c claim) { noOverlap(id, fabric, c) }(v)
	}

	return <-id
}

func readAndFindNoOverlap(r io.Reader) int {
	var claims []claim
	height := 0
	re := regexp.MustCompile(`#(?P<id>\d*) @ (?P<left>\d*),(?P<top>\d*): (?P<width>\d*)x(?P<height>\d*)`)
	scanner := bufio.NewScanner(r)
	width := 0

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		match := re.FindStringSubmatch(txt)
		c := claim{
			unsafeAtoi(match[1]),
			unsafeAtoi(match[2]),
			unsafeAtoi(match[3]),
			unsafeAtoi(match[4]),
			unsafeAtoi(match[5]),
		}
		rc := c.left + c.width
		bl := c.top + c.height

		if rc > width {
			width = rc
		}

		if bl > height {
			height = bl
		}

		claims = append(claims, c)
	}

	return findNoOverlap(width, height, claims)
}

func unsafeAtoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return i
}

func main() {
	fmt.Printf("%v", readAndFindNoOverlap(os.Stdin))
}
