package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type node struct {
	childCount, metaCount, value int
	meta                         []int
	children                     []*node
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
}

func newNode(childCount, metaCount int) *node {
	return &node{childCount, metaCount, 0, []int{}, []*node{}}
}

func unsafeInt(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

type stack struct {
	nodes []*node
}

func (s *stack) add(n *node) {
	s.nodes = append(s.nodes, n)
}

func (s *stack) pop() *node {
	if len(s.nodes) == 0 {
		return nil
	}

	n := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]

	return n
}

func (s *stack) last() (n *node, ok bool) {
	if len(s.nodes) == 0 {
		return nil, false
	}

	return s.nodes[len(s.nodes)-1], true
}

func (s *stack) len() int {
	return len(s.nodes)
}

func makeMeta(metaCount, start int, items []string) (meta []int, sum int) {
	for j := 0; j < metaCount; j++ {
		v := unsafeInt(items[start+j])
		sum += v
		meta = append(meta, v)
	}

	return meta, sum
}

func rootValue(input string) int {
	lns := strings.Split(input, " ")
	parentStack := &stack{}
	parentStack.add(newNode(unsafeInt(lns[0]), unsafeInt(lns[1])))

	lns = lns[2:]
	for i := 0; i < len(lns); i++ {
		childCount := unsafeInt(lns[i])
		metaCount := unsafeInt(lns[i+1])
		i++

		n := newNode(childCount, metaCount)

		currentParent, _ := parentStack.last()
		currentParent.addChild(n)

		if childCount != 0 {
			parentStack.add(n)

			continue
		}

		m, s := makeMeta(metaCount, i+1, lns)
		n.meta = m
		n.value = s
		i += len(m)

		for {
			cp, _ := parentStack.last()

			if len(cp.children) != cp.childCount {
				break
			}

			m, _ := makeMeta(cp.metaCount, i+1, lns)
			cp.meta = m

			for _, indx := range cp.meta {
				if indx <= len(cp.children) {
					cp.value += cp.children[indx-1].value
				}
			}

			i += len(m)

			if parentStack.len() == 1 {
				break
			}

			parentStack.pop()
		}
	}

	v, ok := parentStack.last()

	if !ok {
		return 0
	}

	return v.value
}

func main() {
	pwd, _ := os.Getwd()
	txt, _ := ioutil.ReadFile(filepath.Join(pwd, "input.txt"))

	fmt.Printf("%v\n", rootValue(string(txt)))
}
