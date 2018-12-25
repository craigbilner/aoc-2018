package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"sync"
)

type pot struct {
	value bool
	ll    *pot
	l     *pot
	r     *pot
	rr    *pot
}

type state struct {
	pots     map[int]*pot
	changes  []int
	min, max int
}

func (s *state) addPot(value bool) {
	if len(s.pots) == 0 {
		s.pots[0] = &pot{
			value: value,
		}
		s.pots[-1] = &pot{
			value: false,
			r:     s.pots[0],
		}
		s.pots[-2] = &pot{
			value: false,
			r:     s.pots[-1],
			rr:    s.pots[0],
		}
		s.pots[-3] = &pot{
			value: false,
			r:     s.pots[-2],
			rr:    s.pots[-1],
		}
		s.pots[-4] = &pot{
			value: false,
			r:     s.pots[-3],
			rr:    s.pots[-2],
		}
		s.pots[0].l = s.pots[-1]
		s.pots[0].ll = s.pots[-2]
		s.pots[-1].l = s.pots[-2]
		s.pots[-1].ll = s.pots[-3]
		s.pots[-2].l = s.pots[-3]
		s.pots[-2].ll = s.pots[-4]
		s.pots[-3].l = s.pots[-4]
		s.min = -4
		s.max = 0

		return
	}

	s.pots[s.max+1] = &pot{
		value: value,
		l:     s.pots[s.max],
		ll:    s.pots[s.max-1],
	}
	s.pots[s.max].r = s.pots[s.max+1]
	s.pots[s.max-1].rr = s.pots[s.max+1]

	s.max++

	// fmt.Printf("max %v\n", s.max)
}

func (s *state) pad() {
	if !(!s.pots[s.min].value && !s.pots[s.min+1].value && !s.pots[s.min+2].value && !s.pots[s.min+3].value) {
		s.pots[s.min-1] = &pot{
			value: false,
		}
		s.pots[s.min-2] = &pot{
			value: false,
			r:     s.pots[s.min-1],
		}
		s.pots[s.min-3] = &pot{
			value: false,
			r:     s.pots[s.min-2],
			rr:    s.pots[s.min-1],
		}
		s.pots[s.min-4] = &pot{
			value: false,
			r:     s.pots[-3],
			rr:    s.pots[-2],
		}
		s.pots[s.min+1].ll = s.pots[s.min-1]
		s.pots[s.min].l = s.pots[s.min-1]
		s.pots[s.min-1].l = s.pots[s.min-2]
		s.pots[s.min-1].ll = s.pots[s.min-3]
		s.pots[s.min-2].l = s.pots[s.min-3]
		s.pots[s.min-2].ll = s.pots[s.min-4]
		s.pots[s.min-3].l = s.pots[s.min-4]
		s.min = s.min - 4
	}

	if !(!s.pots[s.max].value && !s.pots[s.max-1].value && !s.pots[s.max-2].value && !s.pots[s.max-3].value) {
		s.pots[s.max+1] = &pot{
			value: false,
			ll:    s.pots[s.max-1],
			l:     s.pots[s.max],
		}
		s.pots[s.max+2] = &pot{
			value: false,
			ll:    s.pots[s.max],
			l:     s.pots[s.max+1],
		}
		s.pots[s.max+3] = &pot{
			value: false,
			ll:    s.pots[s.max+1],
			l:     s.pots[s.max+2],
		}
		s.pots[s.max+4] = &pot{
			value: false,
			ll:    s.pots[s.max+2],
			l:     s.pots[s.max+3],
		}
		s.pots[s.max-1].rr = s.pots[s.max+1]
		s.pots[s.max].r = s.pots[s.max+1]
		s.pots[s.max].rr = s.pots[s.max+2]
		s.pots[s.max+1].r = s.pots[s.max+2]
		s.pots[s.max+1].rr = s.pots[s.max+3]
		s.pots[s.max+2].r = s.pots[s.max+3]
		s.pots[s.max+2].rr = s.pots[s.max+4]
		s.pots[s.max+3].r = s.pots[s.max+4]

		s.max += 4
	}

	// fmt.Printf("padding and setting %v %v\n", s.min, s.max)
}

func (s *state) applyInstructions(grow, noGrow []*instruction) {
	if len(s.changes) == 0 {
		s.generateChanges(grow, noGrow)
		return
	}

	s.generate(grow, noGrow)
}

func (s *state) generateChanges(grow, noGrow []*instruction) {
	// fmt.Printf("generateChanges min %v max %v\n", s.min, s.max)
	for i := s.min + 2; i < s.max-2; i++ {
		// fmt.Printf("i %v\n", i)
		if s.pots[i].value {
			for _, ng := range noGrow {
				if s.pots[i].ll.value == ng.ll && s.pots[i].l.value == ng.l && s.pots[i].value == ng.c && s.pots[i].r.value == ng.r && s.pots[i].rr.value == ng.rr {
					s.changes = append(s.changes, i)
					break
				}
			}

			continue
		}

		for _, g := range grow {
			if s.pots[i].ll.value == g.ll && s.pots[i].l.value == g.l && s.pots[i].value == g.c && s.pots[i].r.value == g.r && s.pots[i].rr.value == g.rr {
				s.changes = append(s.changes, i)
				break
			}
		}
	}

	for _, c := range s.changes {
		s.pots[c].value = !s.pots[c].value
	}
}

type batch struct {
	from, to int
}

func (s *state) generate(grow, noGrow []*instruction) {
	sort.Ints(s.changes)
	// fmt.Printf("generate %v\n", s.changes)

	var batches []*batch
	var b *batch
	for i, c := range s.changes {
		if i == 0 {
			b = &batch{
				c - 2, c + 2,
			}
			continue
		}

		if c <= b.to || (c-1) <= b.to || (c-2) <= b.to {
			b.to = c + 2
			// fmt.Printf("extended %v to %v\n", c, b.to)
			continue
		}

		batches = append(batches, b)
		// fmt.Printf("adding b %v\n", b)
		b = &batch{
			c - 2, c + 2,
		}
	}
	// fmt.Printf("adding b %v\n", b)
	batches = append(batches, b)
	// fmt.Printf("batches %#v\n", batches)

	ch := make(chan []int, len(batches))
	var wg sync.WaitGroup
	for _, bb := range batches {
		wg.Add(1)
		go func(bbb *batch) {
			var changes []int
			for i := bbb.from; i < bbb.to+1; i++ {
				// fmt.Printf("i %v\n", i)
				if s.pots[i].value {
					for _, ng := range noGrow {
						if s.pots[i].ll.value == ng.ll && s.pots[i].l.value == ng.l && s.pots[i].value == ng.c && s.pots[i].r.value == ng.r && s.pots[i].rr.value == ng.rr {
							changes = append(changes, i)
							break
						}
					}

					continue
				}

				for _, g := range grow {
					// fmt.Printf("pot %v %#v\n", i, s.pots[i])
					if s.pots[i].ll.value == g.ll && s.pots[i].l.value == g.l && s.pots[i].value == g.c && s.pots[i].r.value == g.r && s.pots[i].rr.value == g.rr {
						changes = append(changes, i)
						break
					}
				}
			}
			ch <- changes
			wg.Done()
		}(bb)
	}

	wg.Wait()
	close(ch)

	var pots []int
	for ps := range ch {
		pots = append(pots, ps...)
		for _, p := range ps {
			s.pots[p].value = !s.pots[p].value
		}
	}

	s.changes = pots
}

func (s *state) total() int {
	sum := 0

	for k, p := range s.pots {
		if p.value {
			sum += k
		}
	}

	return sum
}

func newState() *state {
	return &state{
		pots:    make(map[int]*pot),
		changes: []int{},
	}
}

type instruction struct {
	ll, l, c, r, rr, result bool
}

func stringToPlant(s string) bool {
	if s == "#" {
		return true
	}

	return false
}

func sumPlantPots(generations int, s *state, grow, noGrow []*instruction) int {
	i := 0
	var gens []int
	for {
		gens = append(gens, s.total())

		if i == generations {
			break
		}

		s.applyInstructions(grow, noGrow)
		s.pad()

		i++
	}

	return s.total()
}

func readAndSumPlantPots(generations int, r *bufio.Reader) int {
	initialStateRe := regexp.MustCompile(`initial state: (.*)`)
	instructionRe := regexp.MustCompile(`([#\.]{5}) => ([#\.])`)
	count := 0
	state := newState()
	var grow []*instruction
	var noGrow []*instruction

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		count++

		if count == 1 {
			stateMatch := initialStateRe.FindStringSubmatch(string(line))

			if len(stateMatch) == 0 {
				return 0
			}

			for _, c := range stateMatch[1] {
				state.addPot(stringToPlant(string(c)))
			}

			continue
		}

		if len(string(line)) == 0 {
			if count > 2 {
				break
			}

			continue
		}

		instMatch := instructionRe.FindStringSubmatch(string(line))

		if len(instMatch) == 0 {
			continue
		}

		seq := instMatch[1]
		result := instMatch[2]
		resultPlant := stringToPlant(result)

		if resultPlant {
			grow = append(grow, &instruction{
				stringToPlant(string(seq[0])),
				stringToPlant(string(seq[1])),
				stringToPlant(string(seq[2])),
				stringToPlant(string(seq[3])),
				stringToPlant(string(seq[4])),
				resultPlant,
			})
			continue
		}

		noGrow = append(noGrow, &instruction{
			stringToPlant(string(seq[0])),
			stringToPlant(string(seq[1])),
			stringToPlant(string(seq[2])),
			stringToPlant(string(seq[3])),
			stringToPlant(string(seq[4])),
			resultPlant,
		})
	}

	state.pad()

	return sumPlantPots(generations, state, grow, noGrow)
}

func main() {
	pwd, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(pwd, "input.txt"))
	r := bufio.NewReader(file)

	fmt.Printf("answer %v\n", readAndSumPlantPots(20, r))
}
