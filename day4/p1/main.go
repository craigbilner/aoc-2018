package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type event int

const (
	start event = iota
	wake
	sleep
)

type item struct {
	evt       event
	timestamp time.Time
	id        int
}

type byTimestamp []item

func (items byTimestamp) Len() int {
	return len(items)
}
func (items byTimestamp) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
func (items byTimestamp) Less(i, j int) bool {
	return items[i].timestamp.Before(items[j].timestamp)
}

type guard struct {
	minutesAsleep     int
	sleepyMinute      int
	sleepyMinuteCount int
	asleep            map[int]int
}

func (g *guard) addSleep(from, to int) {
	for i := from; i < to; i++ {
		g.asleep[i]++

		if g.asleep[i] > g.sleepyMinuteCount {
			g.sleepyMinuteCount = g.asleep[i]
			g.sleepyMinute = i
		}

		g.minutesAsleep++
	}
}

func newGuard() *guard {
	return &guard{
		asleep: make(map[int]int),
	}
}

func guardDetails(items []item) map[int]*guard {
	details := make(map[int]*guard)
	currentId := 0
	var sleepStart time.Time
	for _, i := range items {
		if i.evt == start {
			currentId = i.id

			if _, ok := details[currentId]; !ok {
				details[currentId] = newGuard()
			}

			continue
		}

		if i.evt == sleep {
			sleepStart = i.timestamp
			continue
		}

		if i.evt == wake {
			details[currentId].addSleep(sleepStart.Minute(), i.timestamp.Minute())
		}
	}

	return details
}

func readAndFindSleepyGuard(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	startsShiftRe := regexp.MustCompile(`\[(?P<date>.*)\] Guard #(?P<id>\d*)`)
	fallsAsleepRe := regexp.MustCompile(`\[(?P<date>.*)\] falls asleep`)
	wakesupRe := regexp.MustCompile(`\[(?P<date>.*)\] wakes up`)
	layout := "2006-01-02 15:04"
	var items byTimestamp

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) == 0 {
			break
		}

		match := startsShiftRe.FindStringSubmatch(txt)

		if len(match) > 0 {
			t, _ := time.Parse(layout, match[1])
			id, _ := strconv.Atoi(match[2])

			items = append(items, item{start, t, id})
		}

		match = fallsAsleepRe.FindStringSubmatch(txt)

		if len(match) > 0 {
			t, _ := time.Parse(layout, match[1])

			items = append(items, item{sleep, t, 0})
		}

		match = wakesupRe.FindStringSubmatch(txt)

		if len(match) > 0 {
			t, _ := time.Parse(layout, match[1])

			items = append(items, item{wake, t, 0})
		}
	}

	sort.Sort(items)

	guards := guardDetails(items)

	minutesAsleep := 0
	sleepyId := 0
	sleepyMinute := 0

	for id, guard := range guards {
		if guard.minutesAsleep > minutesAsleep {
			minutesAsleep = guard.minutesAsleep
			sleepyId = id
			sleepyMinute = guard.sleepyMinute
		}
	}

	return sleepyId * sleepyMinute
}

func main() {
	fmt.Printf("%v", readAndFindSleepyGuard(os.Stdin))
}
