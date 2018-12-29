package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

const (
	up int = iota
	right
	down
	left
	straight

	upcart
	downcart
	leftcart
	rightcart

	v
	h
	cr
	cl
	intsect
	blank
)

type coord struct {
	x, y int
}

type rail struct {
	what int
	*coord
	p1, p2, p3, p4 *rail
}

func (r *rail) next(direction, preferredDirection int) (*rail, int) {
	if r.what == v && direction == up {
		return r.p1, up
	}

	if r.what == v && direction == down {
		return r.p2, down
	}

	if r.what == h && direction == left {
		return r.p1, left
	}

	if r.what == h && direction == right {
		return r.p2, right
	}

	if r.what == cr && direction == up {
		return r.p1, right
	}

	if r.what == cr && direction == right {
		return r.p1, up
	}

	if r.what == cr && direction == down {
		return r.p2, left
	}

	if r.what == cr && direction == left {
		return r.p2, down
	}

	if r.what == cl && direction == up {
		return r.p1, left
	}

	if r.what == cl && direction == right {
		return r.p2, down
	}

	if r.what == cl && direction == down {
		return r.p2, right
	}

	if r.what == cl && direction == left {
		return r.p1, up
	}

	if r.what == intsect && preferredDirection == straight && direction == up {
		return r.p1, up
	}

	if r.what == intsect && preferredDirection == straight && direction == right {
		return r.p2, right
	}

	if r.what == intsect && preferredDirection == straight && direction == down {
		return r.p3, down
	}

	if r.what == intsect && preferredDirection == straight && direction == left {
		return r.p4, left
	}

	var exit int
	if r.what == intsect && preferredDirection == left {
		exit = (direction + 3) % 4
	}
	if r.what == intsect && preferredDirection == right {
		exit = (direction + 1) % 4
	}

	if exit == up {
		return r.p1, up
	}

	if exit == right {
		return r.p2, right
	}

	if exit == down {
		return r.p3, down
	}

	if exit == left {
		return r.p4, left
	}

	panic(fmt.Sprintf("Can't find next %#v, %v, %v", r, direction, preferredDirection))
}

type cart struct {
	direction, preferredDirection int
	*rail
}

func (c *cart) move(r *rail, direction int) {
	if c.rail.what == intsect {
		if c.preferredDirection == left {
			c.preferredDirection = straight
		} else if c.preferredDirection == straight {
			c.preferredDirection = right
		} else {
			c.preferredDirection = left
		}
	}

	c.rail = r
	c.direction = direction
}

type cartByCoord []*cart

func (c cartByCoord) Len() int {
	return len(c)
}

func (c cartByCoord) Less(i, j int) bool {
	if c[i].y < c[j].y {
		return true
	}

	if c[i].y > c[j].y {
		return false
	}

	return c[i].x < c[j].x
}

func (c cartByCoord) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func newCart(direction int, r *rail) *cart {
	return &cart{
		direction:          direction,
		rail:               r,
		preferredDirection: left,
	}
}

type track map[string]*rail

func coordToString(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func charToPiece(r rune) int {
	switch r {
	case '|':
		return v
	case '-':
		return h
	case '/':
		return cr
	case '\\':
		return cl
	case '+':
		return intsect
	case '^':
		return upcart
	case '<':
		return leftcart
	case 'v':
		return downcart
	case '>':
		return rightcart
	}

	return blank
}

func neighbours(t track, r *rail) []*rail {
	var rs []*rail

	n, nok := t[coordToString(r.x, r.y-1)]
	e, eok := t[coordToString(r.x+1, r.y)]
	s, sok := t[coordToString(r.x, r.y+1)]
	w, wok := t[coordToString(r.x-1, r.y)]

	if !nok {
		rs = append(rs, nil)
	} else {
		rs = append(rs, n)
	}

	if !eok {
		rs = append(rs, nil)
	} else {
		rs = append(rs, e)
	}

	if !sok {
		rs = append(rs, nil)
	} else {
		rs = append(rs, s)
	}

	if !wok {
		rs = append(rs, nil)
	} else {
		rs = append(rs, w)
	}

	return rs
}

func joinTrack(t track) {
	for _, r := range t {
		ns := neighbours(t, r)

		if r.what == v {
			r.p1 = ns[0]
			r.p2 = ns[2]
			continue
		}

		if r.what == h {
			r.p1 = ns[3]
			r.p2 = ns[1]
			continue
		}

		if r.what == cr && ns[1] != nil && ns[1].what == h {
			r.p1 = ns[1]
			r.p2 = ns[2]
			continue
		}

		if r.what == cr && ns[0] != nil && ns[0].what == v {
			r.p1 = ns[0]
			r.p2 = ns[3]
			continue
		}

		if r.what == cr && ns[1] != nil && ns[1].what == intsect && ns[2] != nil && ns[2].what == v {
			r.p1 = ns[1]
			r.p2 = ns[2]
			continue
		}

		if r.what == cr && ns[2] != nil && ns[2].what == intsect && ns[1] != nil && ns[1].what == h {
			r.p1 = ns[1]
			r.p2 = ns[2]
			continue
		}

		if r.what == cr && ns[0] != nil && ns[0].what == intsect && ns[3] != nil && ns[3].what == h {
			r.p1 = ns[0]
			r.p2 = ns[3]
			continue
		}

		if r.what == cr && ns[3] != nil && ns[3].what == intsect && ns[0] != nil && ns[0].what == v {
			r.p1 = ns[0]
			r.p2 = ns[3]
			continue
		}

		if r.what == cr && ns[2] != nil && ns[2].what == intsect && ns[1] != nil && ns[1].what == intsect {
			r.p1 = ns[1]
			r.p2 = ns[2]
			continue
		}

		if r.what == cr && ns[0] != nil && ns[0].what == intsect && ns[3] != nil && ns[3].what == intsect {
			r.p1 = ns[0]
			r.p2 = ns[3]
			continue
		}

		if r.what == cl && ns[3] != nil && ns[3].what == h {
			r.p1 = ns[3]
			r.p2 = ns[2]
			continue
		}

		if r.what == cl && ns[0] != nil && ns[0].what == v {
			r.p1 = ns[0]
			r.p2 = ns[1]
			continue
		}

		if r.what == cl && ns[3] != nil && ns[3].what == intsect && ns[2] != nil && ns[2].what == v {
			r.p1 = ns[3]
			r.p2 = ns[2]
			continue
		}

		if r.what == cl && ns[2] != nil && ns[2].what == intsect && ns[3] != nil && ns[3].what == h {
			r.p1 = ns[3]
			r.p2 = ns[2]
			continue
		}

		if r.what == cl && ns[0] != nil && ns[0].what == intsect && ns[1] != nil && ns[1].what == h {
			r.p1 = ns[0]
			r.p2 = ns[1]
			continue
		}

		if r.what == cl && ns[1] != nil && ns[1].what == intsect && ns[0] != nil && ns[0].what == v {
			r.p1 = ns[0]
			r.p2 = ns[1]
			continue
		}

		if r.what == cl && ns[3] != nil && ns[3].what == intsect && ns[2] != nil && ns[2].what == intsect {
			r.p1 = ns[3]
			r.p2 = ns[2]
			continue
		}

		if r.what == cl && ns[0] != nil && ns[0].what == intsect && ns[1] != nil && ns[1].what == intsect {
			r.p1 = ns[0]
			r.p2 = ns[1]
			continue
		}

		if r.what == intsect {
			r.p1 = ns[0]
			r.p2 = ns[1]
			r.p3 = ns[2]
			r.p4 = ns[3]
			continue
		}

		panic(fmt.Sprintf("Cannot join %v %v\n", r.what, r.coord))
	}
}

func play(carts []*cart) string {
	positions := make(map[string]*cart)

	for _, c := range carts {
		key := coordToString(c.rail.coord.x, c.rail.coord.y)

		positions[key] = c
	}

	for len(carts) != 1 {
		sort.Sort(cartByCoord(carts))

		for _, c := range carts {
			old := coordToString(c.x, c.y)

			if _, ok := positions[old]; !ok {
				continue
			}

			delete(positions, old)
			c.move(c.rail.next(c.direction, c.preferredDirection))
			key := coordToString(c.rail.coord.x, c.rail.coord.y)

			if _, ok := positions[key]; ok {
				delete(positions, key)
				continue
			}

			positions[key] = c
		}

		var tmp []*cart
		for _, c := range positions {
			tmp = append(tmp, c)
		}
		carts = tmp
	}

	return coordToString(carts[0].x, carts[0].y)
}

func buildAndPlay(r *bufio.Reader) string {
	y := 0
	track := make(track)
	var carts []*cart

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		for x, c := range string(line) {
			xy := coordToString(x, y)
			p := charToPiece(c)

			if p == upcart {
				track[xy] = &rail{what: v, coord: &coord{x, y}}
				carts = append(carts, newCart(up, track[xy]))
				continue
			}
			if p == leftcart {
				track[xy] = &rail{what: h, coord: &coord{x, y}}
				carts = append(carts, newCart(left, track[xy]))
				continue
			}
			if p == downcart {
				track[xy] = &rail{what: v, coord: &coord{x, y}}
				carts = append(carts, newCart(down, track[xy]))
				continue
			}
			if p == rightcart {
				track[xy] = &rail{what: h, coord: &coord{x, y}}
				carts = append(carts, newCart(right, track[xy]))
				continue
			}

			if p != blank {
				track[xy] = &rail{what: p, coord: &coord{x, y}}
			}
		}

		y++
	}

	joinTrack(track)

	return play(carts)
}

func main() {
	pwd, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(pwd, "input.txt"))
	r := bufio.NewReader(file)

	fmt.Printf("answer %v\n", buildAndPlay(r))
}
