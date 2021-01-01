package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	//fmt.Printf("Part2 %d\n", part2())
}

type MinMax [2]int

func part1() int {
	cups, minMax := parse(readInput())

	cur := cups.Front()
	for i := 0; i < 100; i++ {
		cur = play(cups, minMax, cur)
		printList(cups)
	}

	return score(cups)
}

func printList(cups *list.List) {
	nums := make([]int, cups.Len())
	cur := cups.Front()
	var i int
	for cur != nil {
		nums[i] = cur.Value.(int)
		cur = cur.Next()
		i++
	}

	fmt.Println(nums)
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

func play(cups *list.List, minMax MinMax, cur *list.Element) *list.Element {
	next1 := getNext(cups, cur)
	next2 := getNext(cups, next1)
	next3 := getNext(cups, next2)

	removed := map[int]struct{}{
		next1.Value.(int): {},
		next2.Value.(int): {},
		next3.Value.(int): {},
	}

	destination := getDestination(cups, cur.Value.(int)-1, minMax, removed)
	fmt.Println(destination.Value)
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

func getDestination(l *list.List, target int, minMax MinMax, removed map[int]struct{}) *list.Element {
	cur := l.Front()
	if target < minMax[0] {
		return getDestination(l, minMax[1], minMax, removed)
	}

	if _, ok := removed[target]; ok {
		return getDestination(l, target-1, minMax, removed)
	}

	for cur != nil {
		val := cur.Value.(int)
		if val == target {
			return cur
		}
		cur = cur.Next()
	}

	panic("not found")
}
func part2() int {

	return 0
}
func parse(text string) (*list.List, MinMax) {
	l := list.New()
	lowestVal := math.MaxInt32
	maxVal := math.MinInt32
	for _, v := range strings.Split(text, "") {
		val := atoi(v)
		lowestVal = min(lowestVal, val)
		maxVal = max(maxVal, val)
		l.PushBack(val)
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
