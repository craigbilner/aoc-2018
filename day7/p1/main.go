package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
)

type node struct {
	value   string
	visited bool
	deps    []*node
}

type byValue []*node

func (bv byValue) Len() int {
	return len(bv)
}

func (bv byValue) Less(i, j int) bool {
	return bv[i].value < bv[j].value
}

func (bv byValue) Swap(i, j int) {
	bv[i], bv[j] = bv[j], bv[i]
}

func (n *node) isReady() bool {
	if len(n.deps) == 0 {
		return true
	}

	isReady := true
	for _, d := range n.deps {
		if !d.visited {
			isReady = false
			break
		}
	}

	return isReady
}

func orderSteps(nodes map[string]*node) string {
	steps := ""
	var ready byValue

	for {
		for _, node := range nodes {
			if !node.visited && node.isReady() {
				ready = append(ready, node)
			}
		}

		if len(ready) == 0 {
			break
		}

		sort.Sort(ready)

		steps += ready[0].value
		ready[0].visited = true

		ready = byValue{}
	}

	return steps
}

func readAndOrderSteps(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	re := regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z])`)
	nodes := make(map[string]*node)

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		match := re.FindStringSubmatch(txt)

		if len(match) > 0 {
			_, ok := nodes[match[1]]

			if !ok {
				nodes[match[1]] = &node{value: match[1], deps: []*node{}}
			}

			_, ok1 := nodes[match[2]]

			if !ok1 {
				nodes[match[2]] = &node{value: match[2], deps: []*node{}}
			}

			nodes[match[2]].deps = append(nodes[match[2]].deps, nodes[match[1]])
		}
	}

	return orderSteps(nodes)
}

func main() {
	fmt.Printf("%v\n", readAndOrderSteps(os.Stdin))
}
