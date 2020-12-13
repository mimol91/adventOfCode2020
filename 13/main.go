package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	text := readInput()
	minDepartTime, buses := parse(text)
	var chosenBus int
	minWaitTime := math.MaxInt32

	for _, bus := range buses {
		val := bus * ((minDepartTime / bus) + 1)
		diff := val - minDepartTime
		if diff < minWaitTime {
			minWaitTime = diff
			chosenBus = bus
		}
	}

	return chosenBus * minWaitTime
}

func parse(text string) (int, []int) {
	lines := strings.Split(text, "\n")
	minDepartTime := atoi(lines[0])
	vals := strings.Split(lines[1], ",")
	buses := make([]int, 0)
	for _, val := range vals {
		if val != "x" {
			buses = append(buses, atoi(val))
		}
	}
	return minDepartTime, buses
}

func part2() int {
	text := readInput()
	lines := strings.Split(text, "\n")
	buses := make([]int, 0)
	for _, val := range strings.Split(lines[1], ",") {
		if val == "x" {
			buses = append(buses, 0)
		} else {
			buses = append(buses, atoi(val))
		}
	}

	var start int
	interval := buses[0]

	for i, currentBus := range buses[1:] {
		if currentBus == 0 {
			continue
		}
		for (start+i+1)%currentBus != 0 {
			start += interval
		}
		interval *= currentBus
	}

	return start
}

func readInput() string {
	b, err := ioutil.ReadFile("13/input.txt")
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
