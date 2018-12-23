package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func printPoints(f *os.File, s *sky, points map[int][]int) {
	for y := s.topLeft.y; y < s.bottomRight.y+1; y++ {
		line := make([]int, s.bottomRight.x-s.topLeft.x+1)

		v, ok := points[y]
		if !ok {
			continue
		}

		for _, px := range v {
			line[px-s.topLeft.x] = 1
		}

		f.WriteString(fmt.Sprintf("%v\n", line))
	}
}

func calculatePosition(points []*point, t int) (*sky, map[int][]int) {
	m := make(map[int][]int)
	left, right, up, down := 0, 0, 0, 0

	for _, p := range points {
		x := p.position.x + (p.velocity.x * t)
		y := p.position.y + (p.velocity.y * t)

		if _, ok := m[y]; !ok {
			m[y] = []int{x}
			continue
		}

		m[y] = append(m[y], x)

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
	}, m
}

func readAndPrint(from, duration int, r io.Reader) {
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

	pwd, _ := os.Getwd()

	for i := from; i < from + duration; i++ {
		f, err := os.Create(filepath.Join(pwd, fmt.Sprintf("output_%v.txt", i)))
		defer f.Close()

		if err != nil {
			println(err)
			return
		}

		s, pointsByRow := calculatePosition(points, i)
		f.WriteString(fmt.Sprintf("Time: %v\n", i))
		printPoints(f, s, pointsByRow)
	}
}

func main() {
	readAndPrint(10000, 200, os.Stdin)
}
