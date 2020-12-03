package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part1: %s \n\nPart2: %s", part1(), part2())
}

func part1() string {
	data := readInput()
	list := strings.Split(data, "\n")
	nums := make([]int, len(list))
	for i, row := range list {
		num, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	mem := map[int]struct{}{}
	for _, num := range nums {
		reaming := 2020 - num
		if _, ok := mem[num]; ok {
			return strconv.Itoa(num * reaming)
		} else {
			mem[reaming] = struct{}{}
		}
	}

	return ""
}
func part2() string {

	data := readInput()
	list := strings.Split(data, "\n")
	nums := make([]int, len(list))
	for i, row := range list {
		num, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	mem := map[int]int{}
	for _, num := range nums {
		reaming := 2020 - num
		mem[reaming] = 0
	}
	for val, _ := range mem {
		for _, num := range nums {
			reaming := val - num
			mem[reaming] = num
		}
	}
	for val, c := range mem {
		for _, num := range nums {
			if val == num && c != 0 {
				g := 2020 - (val + c)
				return strconv.Itoa(val * c * g)
			}
		}
	}
	return ""
}

func readInput() string {
	b, err := ioutil.ReadFile("01/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
