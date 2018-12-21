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
	childCount, metaCount int
	meta                  []int
	children              []*node
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
}

func newNode(childCount, metaCount int) *node {
	return &node{childCount, metaCount, []int{}, []*node{}}
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

func sumMeta(input string) int {
	lns := strings.Split(input, " ")
	parentStack := &stack{}
	parentStack.add(newNode(unsafeInt(lns[0]), unsafeInt(lns[1])))

	lns = lns[2:]
	sum := 0
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
		sum += s
		i += len(m)

		for {
			cp, _ := parentStack.last()

			if len(cp.children) != cp.childCount {
				break
			}

			m, s := makeMeta(cp.metaCount, i+1, lns)
			cp.meta = m
			sum += s
			i += len(m)

			if parentStack.len() == 1 {
				break
			}

			parentStack.pop()
		}
	}

	return sum
}

func main() {
	pwd, _ := os.Getwd()
	txt, _ := ioutil.ReadFile(filepath.Join(pwd, "input.txt"))

	fmt.Printf("%v\n", sumMeta(string(txt)))
}
