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

func part1() int {
	return execute(2020)
}

func part2() int {
	return execute(30000000)
}

func execute(maxIterations int) int {
	text := readInput()
	startingNumbers := parse(text)
	occurrencesMap := make([]int, maxIterations)

	for i, val := range startingNumbers {
		occurrencesMap[val] = i + 1
	}

	lastSpoken := startingNumbers[len(startingNumbers)-1]
	var val int
	for i := len(startingNumbers); i < maxIterations; i++ {
		val = occurrencesMap[lastSpoken]
		if val == 0 {
			occurrencesMap[lastSpoken] = i
			lastSpoken = 0
		} else {
			tmp := lastSpoken
			lastSpoken = i - val
			occurrencesMap[tmp] = i
		}
	}

	return lastSpoken
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
