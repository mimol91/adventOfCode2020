package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {

	return 0
}
func part2() int {

	return 0
}

func readInput() string {
	b, err := ioutil.ReadFile("01/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
