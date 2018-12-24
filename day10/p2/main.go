package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type point struct {
	position, velocity *coord
}

type sky struct {
	topLeft, bottomRight *coord
}

func unsafeAtoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return i
}

func calculatePosition(points []*point, t int) *sky {
	left, right, up, down := 0, 0, 0, 0

	for _, p := range points {
		x := p.position.x + (p.velocity.x * t)
		y := p.position.y + (p.velocity.y * t)

		if x < left {
			left = x
		}

		if y < up {
			up = y
		}

		if x > right {
			right = x
		}

		if y > down {
			down = y
		}
	}

	return &sky{
		&coord{
			left,
			up,
		},
		&coord{
			right,
			down,
		},
	}
}

func readAndTime(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	re := regexp.MustCompile(`position=<([0-9\-\s]*), ([0-9\-\s]*)> velocity=<([0-9\-\s]*), ([0-9\-\s]*)>`)
	var points []*point

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		match := re.FindStringSubmatch(txt)

		if len(match) == 0 {
			continue
		}

		p := &point{
			&coord{
				unsafeAtoi(strings.Trim(match[1], " ")),
				unsafeAtoi(strings.Trim(match[2], " ")),
			},
			&coord{
				unsafeAtoi(strings.Trim(match[3], " ")),
				unsafeAtoi(strings.Trim(match[4], " ")),
			},
		}
		points = append(points, p)
	}

	t := 0
	s := calculatePosition(points, 0)
	for {
		t++

		news := calculatePosition(points, t)

		if news.topLeft.x < s.topLeft.x ||
			news.topLeft.y < s.topLeft.y ||
			news.bottomRight.x > s.bottomRight.x ||
			news.bottomRight.y > s.bottomRight.y {
			break
		}

		s = news
	}

	return t - 1
}

func main() {
	fmt.Printf("%v\n", readAndTime(os.Stdin))
}
