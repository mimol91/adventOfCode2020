package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

type Tuple struct {
	A, B int
}

func (r Tuple) Diff() int { return r.B - r.A }
func (r *Tuple) Store(val int) {
	r.A = r.B
	r.B = val
}
func part1() int {
	return execute(2020)
}

func part2() int {
	return execute(30000000)
}

func execute(maxIterations int) int {
	text := readInput()
	startingNumbers := parse(text)
	occurrencesMap := make(map[int]*Tuple)

	for i, val := range startingNumbers {
		occurrencesMap[val] = &Tuple{B: i + 1}
	}

	lastNum := startingNumbers[len(startingNumbers)-1]
	isNew := true
	for i := 1 + len(occurrencesMap); i <= maxIterations; i++ {
		if isNew {
			lastNum = 0
		} else {
			lastNum = occurrencesMap[lastNum].Diff()
		}
		_, ok := occurrencesMap[lastNum]
		if !ok {
			occurrencesMap[lastNum] = &Tuple{}
		}
		isNew = !ok
		occurrencesMap[lastNum].Store(i)
	}

	return lastNum
}

func parse(text string) []int {
	elements := strings.Split(text, ",")
	result := make([]int, len(elements))
	for i, v := range elements {
		result[i] = atoi(v)
	}

	return result
}

func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}

func readInput() string {
	b, err := ioutil.ReadFile("15/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
