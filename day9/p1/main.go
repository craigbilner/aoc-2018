package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type circle struct {
	currentMarbleIndx int
	marbles           []int
}

func (c *circle) add(marble int) {
	newIndx := (c.currentMarbleIndx + 2) % len(c.marbles)

	if newIndx == 0 {
		c.marbles = append(c.marbles, marble)
		c.currentMarbleIndx = len(c.marbles) - 1
		return
	}

	c.marbles = append(c.marbles, 0)
	copy(c.marbles[newIndx+1:], c.marbles[newIndx:])
	c.marbles[newIndx] = marble
	c.currentMarbleIndx = newIndx
}

func (c *circle) remove() int {
	removeIndx := c.currentMarbleIndx - 7
	if removeIndx < 0 {
		removeIndx = len(c.marbles) + removeIndx
	}

	removed := c.marbles[removeIndx]
	c.marbles = append(c.marbles[:removeIndx], c.marbles[removeIndx+1:]...)
	c.currentMarbleIndx = removeIndx

	return removed
}

func newCircle() *circle {
	return &circle{
		0,
		[]int{0},
	}
}

type player struct {
	score int
}

func unsafeAtoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return i
}

func highScore(players, lastMarble int) int {
	c := newCircle()

	var ps []*player
	for i := 0; i < players; i++ {
		ps = append(ps, &player{})
	}

	currentPlayer := 1
	currentMarble := 1

	for {
		if currentMarble%23 == 0 {
			removed := c.remove()
			ps[currentPlayer-1].score += currentMarble + removed
		} else {
			c.add(currentMarble)
		}

		if currentMarble == lastMarble {
			break
		}

		currentMarble++

		if currentPlayer == players {
			currentPlayer = 0
		}

		currentPlayer++
	}

	hs := 0
	for _, p := range ps {
		if p.score > hs {
			hs = p.score
		}
	}

	return hs
}

func highScoreFromText(s string) int {
	re := regexp.MustCompile(`(\d*) players; last marble is worth (\d*) points`)
	match := re.FindStringSubmatch(s)

	if len(match) == 0 {
		return 0
	}

	return highScore(unsafeAtoi(match[1]), unsafeAtoi(match[2]))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	println(highScoreFromText(text))
}
