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
	value    string
	visited  bool
	visiting bool
	deps     []*node
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

func stepToTime(value string) int {
	return int(value[0] - 64)
}

type worker struct {
	isWorking bool
	workingOn *node
	timeLeft  int
}

func (w *worker) startWork(n *node, stepTime int) {
	w.workingOn = n
	w.workingOn.visiting = true
	w.timeLeft = stepToTime(n.value) + stepTime
	w.isWorking = true
}

func (w *worker) decrement() (finished bool) {
	w.timeLeft--

	if w.timeLeft == 0 {
		w.isWorking = false
		w.workingOn.visiting = false
		w.workingOn.visited = true
		return true
	}

	return false
}

func whoIsReady(nodes map[string]*node) byValue {
	var ready byValue
	for _, node := range nodes {
		if !node.visited && !node.visiting && node.isReady() {
			ready = append(ready, node)
		}
	}

	sort.Sort(ready)

	return ready
}

func timeSteps(stepTime, workerCount, nodeCount int, nodes map[string]*node) int {
	var workers []*worker

	for i := 0; i < workerCount; i++ {
		workers = append(workers, &worker{})
	}

	visitedCount := 0
	time := 0
	for {
		if visitedCount == nodeCount {
			break
		}

		ready := whoIsReady(nodes)
		for _, w := range workers {
			if !w.isWorking && len(ready) > 0 {
				w.startWork(ready[0], stepTime)
				ready = ready[1:]
			}

			finished := w.decrement()

			if finished {
				visitedCount++
			}
		}

		time++
	}

	return time
}

func readAndOTimeSteps(stepTime, workerCount int, r io.Reader) int {
	scanner := bufio.NewScanner(r)
	re := regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z])`)
	nodes := make(map[string]*node)
	nodeCount := 0

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
				nodeCount++
			}

			_, ok1 := nodes[match[2]]

			if !ok1 {
				nodes[match[2]] = &node{value: match[2], deps: []*node{}}
				nodeCount++
			}

			nodes[match[2]].deps = append(nodes[match[2]].deps, nodes[match[1]])
		}
	}

	return timeSteps(stepTime, workerCount, nodeCount, nodes)
}

func main() {
	fmt.Printf("%v\n", readAndOTimeSteps(60, 5, os.Stdin))
}
