package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func part2() int {
	text := readInput()
	lines := strings.Split(text, "\n")
	adapters := make([]int, len(lines))
	for i, line := range lines {
		adapters[i] = atoi(line)
	}
	sort.Ints(adapters)
	var res [][2]int

	for _, val := range adapters {
		for i := len(res) - 1; i >= 0; i-- {
			elements := res[i]
			diff := val - elements[1]
			if diff > 3 {
				continue
			}
			res = append(res, [2]int{elements[0], val})
		}
		if val < 4 {
			res = append(res, [2]int{val, val})
		}
	}

	var result int
	lastElement := adapters[len(adapters)-1]
	for _, list := range res {
		lastItem := list[len(list)-1]
		if lastItem == lastElement {
			result++
		}
	}
	return result
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")
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

func BenchmarkAA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2()
	}
}
