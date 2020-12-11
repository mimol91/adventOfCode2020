package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	text := readInput()
	lines := strings.Split(text, "\n")
	adapters := make([]int, len(lines))
	for i, line := range lines {
		adapters[i] = atoi(line)
	}
	sort.Ints(adapters)
	histogram := map[int]int{
		adapters[0]: 1,
	}

	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		histogram[diff]++
	}
	histogram[3]++

	return histogram[1] * histogram[3]
}

func part2() int {
	text := readInput()
	lines := strings.Split(text, "\n")
	adapters := make([]int, len(lines))
	for i, line := range lines {
		adapters[i] = atoi(line)
	}
	sort.Ints(adapters)
	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	ans := make([]int, len(adapters))
	ans[0] = 1

	for i := 1; i < len(ans); i++ {
		if i >= 4 && adapters[i]-adapters[i-4] == 4 {
			ans[i] = ans[i-4] * 7
			continue
		}
		if i >= 2 && adapters[i]-adapters[i-2] == 2 {
			ans[i] = ans[i-1] * 2
			continue
		}
		ans[i] = ans[i-1]
	}
	return ans[len(ans)-1]
}

func readInput() string {
	b, err := ioutil.ReadFile("10/input.txt")
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
