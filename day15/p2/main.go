package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type coord struct {
	x, y int
}

func (c coord) before(t *coord) bool {
	if c.y < t.y {
		return true
	}

	if c.y > t.y {
		return false
	}

	return c.x < t.x
}

func (c coord) String() string {
	return coordToString(c.x, c.y)
}

type coordByReading []*coord

func (cs coordByReading) Len() int {
	return len(cs)
}

func (cs coordByReading) Less(i, j int) bool {
	if cs[i].y < cs[j].y {
		return true
	}

	if cs[i].y > cs[j].y {
		return false
	}

	return cs[i].x < cs[j].x
}

func (cs coordByReading) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

type arena map[string]*cavern

type cavern struct {
	*coord
	distance   map[string]int
	neighbours []*cavern
}

func (c *cavern) String() string {
	return fmt.Sprintf("%v,%v - distance %v", c.x, c.y, c.distance)
}

func newCavern(x, y int) *cavern {
	return &cavern{&coord{x, y}, make(map[string]int), []*cavern{}}
}

type creature int

const (
	_ creature = iota
	elfy
	gobliny
)

type unit struct {
	typ creature
	*coord
	hitPoints int
	power     int
}

func (u *unit) alive() bool {
	return u.hitPoints > 0
}

func (u *unit) attack(power int) bool {
	if u.alive() {
		u.hitPoints -= power
	}

	return u.alive()
}

func (u *unit) move(c *coord) {
	u.coord = c
}

func (u *unit) String() string {
	return fmt.Sprintf("\n%v: %v\n", u.coord, u.hitPoints)
}

func newGoblin(x, y int) *unit {
	return &unit{typ: gobliny, coord: &coord{x, y}, hitPoints: 200, power: 3}
}

func newElf(x, y, power int) *unit {
	return &unit{typ: elfy, coord: &coord{x, y}, hitPoints: 200, power: power}
}

type unitByCoord []*unit

func (c unitByCoord) Len() int {
	return len(c)
}

func (c unitByCoord) Less(i, j int) bool {
	if c[i].y < c[j].y {
		return true
	}

	if c[i].y > c[j].y {
		return false
	}

	return c[i].x < c[j].x
}

func (c unitByCoord) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type target struct {
	enemy *unit
	space *coord
}

func coordToString(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func firstSteps(round int, a arena, positions map[string]*unit, from, to *coord) (int, coordByReading, bool) {
	fstr := from.String()
	did := fmt.Sprintf("%v:%v", round, fstr)
	currentNode := a[fstr]
	currentNode.distance[did] = 0
	var queue []*cavern

	for _, n := range currentNode.neighbours {
		if _, ok := positions[n.coord.String()]; ok {
			continue
		}

		if n.x == to.x && n.y == to.y {
			return 1, []*coord{to}, true
		}

		n.distance[did] = 1
		queue = append(queue, n)
	}

	maxDist := 0
	for len(queue) > 0 {
		for _, c := range queue {
			for _, n := range c.neighbours {
				if _, ok := positions[n.coord.String()]; ok {
					continue
				}

				nd := c.distance[did] + 1

				if maxDist != 0 && nd > maxDist {
					continue
				}

				if d, ok := n.distance[did]; ok && d <= nd {
					continue
				}

				if n.x == to.x && n.y == to.y {
					if maxDist == 0 || nd < maxDist {
						maxDist = nd
					}
				}

				n.distance[did] = nd
				queue = append(queue, n)
			}

			queue = queue[1:]
		}
	}

	sd := a[to.String()].distance[did]
	routes := map[string]*cavern{to.String(): a[to.String()]}
	d := sd

	for {
		var tmp map[string]*cavern

		for _, r := range routes {
			for _, n := range r.neighbours {
				if v, ok := n.distance[did]; ok && v == d {
					tmp[n.coord.String()] = n
					continue
				}

				if v, ok := n.distance[did]; ok && v < d {
					d = n.distance[did]
					tmp = map[string]*cavern{n.coord.String(): n}
				}
			}
		}

		routes = tmp

		if len(routes) == 0 {
			break
		}

		if d == 1 {
			break
		}
	}

	if len(routes) == 0 {
		return 0, []*coord{}, false
	}

	var fs []*coord
	for _, c := range routes {
		fs = append(fs, c.coord)
	}

	sort.Sort(coordByReading(fs))

	return sd, fs, true
}

func nearest(round int, a arena, positions map[string]*unit, from *coord, enemies map[string]*unit) (*coord, bool) {
	var targets []*target
	for _, e := range enemies {
		north := coordToString(e.x, e.y-1)
		if _, aok := a[north]; aok {
			if _, ok := positions[north]; !ok {
				targets = append(targets, &target{e, &coord{e.x, e.y - 1}})
			}
		}

		east := coordToString(e.x+1, e.y)
		if _, aok := a[east]; aok {
			if _, ok := positions[east]; !ok {
				targets = append(targets, &target{e, &coord{e.x + 1, e.y}})
			}
		}

		south := coordToString(e.x, e.y+1)
		if _, aok := a[south]; aok {
			if _, ok := positions[south]; !ok {
				targets = append(targets, &target{e, &coord{e.x, e.y + 1}})
			}
		}

		west := coordToString(e.x-1, e.y)
		if _, aok := a[west]; aok {
			if _, ok := positions[west]; !ok {
				targets = append(targets, &target{e, &coord{e.x - 1, e.y}})
			}
		}
	}

	if len(targets) == 0 {
		return nil, false
	}

	minDistance := 0
	var step *coord
	var destination *coord
	for _, t := range targets {
		d, fs, ok := firstSteps(round, a, positions, from, t.space)

		if !ok {
			continue
		}

		if minDistance == 0 || d < minDistance {
			minDistance = d
			step = fs[0]
			destination = t.space
			continue
		}

		if d == minDistance && t.space.before(destination) {
			step = fs[0]
			destination = t.space
		}
	}

	if step == nil {
		return nil, false
	}

	return step, true
}

func opponent(from *coord, enemies map[string]*unit) (enemy *unit, canAttack bool) {
	var es []*unit

	if v, ok := enemies[coordToString(from.x, from.y-1)]; ok {
		es = append(es, v)
	}

	if v, ok := enemies[coordToString(from.x-1, from.y)]; ok {
		es = append(es, v)
	}

	if v, ok := enemies[coordToString(from.x+1, from.y)]; ok {
		es = append(es, v)
	}

	if v, ok := enemies[coordToString(from.x, from.y+1)]; ok {
		es = append(es, v)
	}

	if len(es) == 0 {
		return nil, false
	}

	n := es[0]
	for _, e := range es {
		if e.hitPoints < n.hitPoints {
			n = e
		}
	}

	return n, true
}

func combat(a arena, elves []*unit, goblins []*unit) (sum int, remainingElves int) {
	round := 0

Loop:
	for {
		var us []*unit
		positions := make(map[string]*unit)
		gmap := make(map[string]*unit)
		emap := make(map[string]*unit)
		for _, v := range elves {
			if !v.alive() {
				continue
			}

			positions[v.coord.String()] = v
			emap[v.coord.String()] = v
			us = append(us, v)
		}
		for _, v := range goblins {
			if !v.alive() {
				continue
			}

			positions[v.coord.String()] = v
			gmap[v.coord.String()] = v
			us = append(us, v)
		}

		sort.Sort(unitByCoord(us))
		for _, u := range us {
			if !u.alive() {
				continue
			}

			if len(emap) == 0 || len(gmap) == 0 {
				break Loop
			}

			var enemy *unit
			hasOpponent := false

			if u.typ == elfy {
				enemy, hasOpponent = opponent(u.coord, gmap)
			} else if u.typ == gobliny {
				enemy, hasOpponent = opponent(u.coord, emap)
			}

			if hasOpponent {
				stillAlive := enemy.attack(u.power)

				if !stillAlive {
					delete(positions, enemy.coord.String())

					if enemy.typ == elfy {
						delete(emap, enemy.coord.String())
					} else if enemy.typ == gobliny {
						delete(gmap, enemy.coord.String())
					}
				}

				continue
			}

			var step *coord
			canReach := false

			if u.typ == elfy {
				step, canReach = nearest(round, a, positions, u.coord, gmap)
			} else if u.typ == gobliny {
				step, canReach = nearest(round, a, positions, u.coord, emap)
			}

			if canReach {
				delete(positions, u.coord.String())

				if u.typ == elfy {
					delete(emap, u.coord.String())
				} else if u.typ == gobliny {
					delete(gmap, u.coord.String())
				}

				u.move(step)

				positions[u.coord.String()] = u

				if u.typ == elfy {
					emap[u.coord.String()] = u
					enemy, hasOpponent = opponent(u.coord, gmap)
				} else if u.typ == gobliny {
					gmap[u.coord.String()] = u
					enemy, hasOpponent = opponent(u.coord, emap)
				}

				if hasOpponent {
					stillAlive := enemy.attack(u.power)

					if !stillAlive {
						delete(positions, enemy.coord.String())

						if enemy.typ == elfy {
							delete(emap, enemy.coord.String())
						} else if enemy.typ == gobliny {
							delete(gmap, enemy.coord.String())
						}
					}
				}
			}
		}

		round++
	}

	for _, u := range elves {
		if u.alive() {
			remainingElves += 1
			sum += u.hitPoints
		}
	}

	sum *= round

	return
}

func fight(a arena, elves []*unit, goblins []*unit) int {
	need := len(elves)
	power := 4
	prev := 0
	max := 4
	min := 4
	sum := 0

	for {
		testArena := make(arena)
		for k, v := range a {
			testArena[k] = newCavern(v.x, v.y)
		}
		makeNeighbours(testArena)
		var testElves []*unit
		for _, e := range elves {
			testElves = append(testElves, newElf(e.x, e.y, power))
		}
		var testGoblins []*unit
		for _, g := range goblins {
			testGoblins = append(testGoblins, newGoblin(g.x, g.y))
		}

		s, remaining := combat(testArena, testElves, testGoblins)

		// fmt.Printf("power %v, prev %v, s: %v, r: %v\n", power, prev, s, remaining)

		if remaining == need && (power == 4 || (prev-power == 1) || (power-prev == 1)) {
			sum = s
			break
		}

		if (prev-power == 1) || (power-prev == 1) {
			break
		}

		if power > max {
			max = power
		}

		if remaining == need && power > prev {
			min = prev
			p := power
			power = prev + ((power - prev) / 2)
			prev = p
			sum = s

			continue
		}

		if remaining == need {
			p := power
			power = min + ((power - min) / 2)
			prev = p
			sum = s

			continue
		}

		if sum > 0 {
			min = power
			p := power
			power = power + ((max - power) / 2)
			prev = p

			continue
		}

		min = prev
		prev = power
		power *= 2
	}

	return sum
}

func makeNeighbours(a arena) {
	for _, v := range a {
		if nc, ok := a[coordToString(v.x, v.y-1)]; ok {
			v.neighbours = append(v.neighbours, nc)
		}
		if nc, ok := a[coordToString(v.x+1, v.y)]; ok {
			v.neighbours = append(v.neighbours, nc)
		}
		if nc, ok := a[coordToString(v.x, v.y+1)]; ok {
			v.neighbours = append(v.neighbours, nc)
		}
		if nc, ok := a[coordToString(v.x-1, v.y)]; ok {
			v.neighbours = append(v.neighbours, nc)
		}
	}
}

func readerToArena(r io.Reader) (arena, []*unit, []*unit) {
	scanner := bufio.NewScanner(r)
	a := make(map[string]*cavern)
	var elves []*unit
	var goblins []*unit
	y := 0

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		x := 0
		for _, c := range txt {
			switch c {
			case '#':
				// do nothing
			case '.':
				a[coordToString(x, y)] = newCavern(x, y)
			case 'G':
				a[coordToString(x, y)] = newCavern(x, y)
				goblins = append(goblins, newGoblin(x, y))
			case 'E':
				a[coordToString(x, y)] = newCavern(x, y)
				elves = append(elves, newElf(x, y, 0))
			default:
				panic(fmt.Sprintf("Unexpected input %v", c))
			}

			x++
		}

		y++
	}

	makeNeighbours(a)

	return a, elves, goblins
}

func main() {
	println(fight(readerToArena(os.Stdin)))
}
