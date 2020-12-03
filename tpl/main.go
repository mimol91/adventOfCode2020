package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part1 %s\n", part1())
	fmt.Printf("Part2 %s\n", part2())

}

func part1() string {

	return ""
}
func part2() string {

	return ""
}

func readInput() string {
	b, err := ioutil.ReadFile("01/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
