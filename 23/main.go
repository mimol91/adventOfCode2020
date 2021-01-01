package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type MinMax [2]int
type Store map[int]*list.Element

const (
	maxCups  = 1000000
	maxMoves = 10000000
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2()) //112	//389125467
}

func part1() int {
	store := make(Store, maxCups)
	cups, minMax := parse(readInput(), store)
	cur := cups.Front()

	for i := 0; i < 100; i++ {
		cur = play(cups, minMax, cur, store)
	}

	return score(cups)
}

func score(cups *list.List) int {
	var score string
	var afterOne bool

	cur := cups.Front()
	for cur != nil {
		val := cur.Value.(int)

		if afterOne {
			score += strconv.Itoa(val)
		} else {
			if val == 1 {
				afterOne = true
			}
		}
		cur = cur.Next()
	}
	cur = cups.Front()
	for cur != nil {
		val := cur.Value.(int)
		if val == 1 {
			break
		}
		score += strconv.Itoa(val)
		cur = cur.Next()
	}

	return atoi(score)
}

func play(cups *list.List, minMax MinMax, cur *list.Element, store Store) *list.Element {
	next1 := getNext(cups, cur)
	next2 := getNext(cups, next1)
	next3 := getNext(cups, next2)

	removed := map[int]struct{}{
		next1.Value.(int): {},
		next2.Value.(int): {},
		next3.Value.(int): {},
	}

	destination := getDestination(cups, cur.Value.(int)-1, minMax, removed, store)

	cups.MoveAfter(next3, destination)
	cups.MoveAfter(next2, destination)
	cups.MoveAfter(next1, destination)

	return getNext(cups, cur)
}

func getNext(cups *list.List, cur *list.Element) *list.Element {
	next := cur.Next()
	if next != nil {
		return next
	}
	return cups.Front()
}

func getDestination(l *list.List, target int, minMax MinMax, removed map[int]struct{}, store Store) *list.Element {
	if target < minMax[0] {
		return getDestination(l, minMax[1], minMax, removed, store)
	}
	if _, ok := removed[target]; ok {
		return getDestination(l, target-1, minMax, removed, store)
	}
	return store[target]
}
func part2() int {
	store := make(Store, maxCups)

	cups, minMax := parse(readInput(), store)
	for i := minMax[1] + 1; i <= maxCups; i++ {
		el := cups.PushBack(i)
		store[i] = el
	}
	minMax[1] = maxCups
	cur := cups.Front()

	for i := 0; i < maxMoves; i++ {
		cur = play(cups, minMax, cur, store)
	}

	return score2(store)
}
func score2(store Store) int {
	one := store[1]

	score := one.Next().Value.(int)
	score *= one.Next().Next().Value.(int)

	return score
}

func parse(text string, store Store) (*list.List, MinMax) {
	l := list.New()
	lowestVal := math.MaxInt32
	maxVal := math.MinInt32
	for _, v := range strings.Split(text, "") {
		val := atoi(v)
		lowestVal = min(lowestVal, val)
		maxVal = max(maxVal, val)
		el := l.PushBack(val)
		store[val] = el
	}
	return l, MinMax{lowestVal, maxVal}
}

func max(val int, val2 int) int {
	if val > val2 {
		return val
	}
	return val2
}

func min(val int, val2 int) int {
	if val < val2 {
		return val
	}
	return val2
}

func readInput() string {
	b, err := ioutil.ReadFile("23/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}
