package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Q []string
type Memory map[int]int

func (r Memory) Sum() int {
	var result int
	for _, val := range r {
		result += val
	}
	return result
}

func (r *Q) Empty() bool      { return len(*r) == 0 }
func (r *Q) Enqueue(p string) { *r = append(*r, p) }
func (r *Q) Dequeue() string {
	if len(*r) == 0 {
		panic("opsie, Q is empty")
	}
	el := (*r)[0]
	*r = (*r)[1:]
	return el
}
func (r *Q) Drain() []string {
	result := make([]string, 0)
	for !r.Empty() {
		result = append(result, r.Dequeue())
	}

	return result
}
func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	text := readInput()
	lines := strings.Split(text, "\n")

	var mask string
	var memVal int
	memory := make(Memory)

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			if _, err := fmt.Sscanf(parts[0], "mem[%d]", &memVal); err != nil {
				panic("unable to parse")
			}
			memory[memVal] = applyMask(atoi(parts[1]), mask)
		}
	}

	return memory.Sum()
}

func part2() int {
	text := readInput()
	lines := strings.Split(text, "\n")

	var mask string
	var memVal int
	memory := make(Memory)

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			if _, err := fmt.Sscanf(parts[0], "mem[%d]", &memVal); err != nil {
				panic("unable to parse")
			}
			for _, address := range generateAddresses(fmt.Sprintf("%036b", memVal), mask) {
				memory[address] = atoi(parts[1])
			}
		}
	}

	return memory.Sum()
}

func generateAddresses(address string, mask string) []int {
	q := Q{address}
	for i, c := range mask {
		if c == '0' {
			continue
		}
		if c == 'X' {
			for _, g := range q.Drain() {
				chars := []rune(g)
				chars[i] = '1'
				q.Enqueue(string(chars))
				chars[i] = '0'
				q.Enqueue(string(chars))
			}
		} else {
			for _, g := range q.Drain() {
				chars := []rune(g)
				chars[i] = c
				q.Enqueue(string(chars))
			}
		}
	}

	list := q.Drain()
	result := make([]int, len(list))
	for i, element := range list {
		result[i] = binToDec(element)
	}

	return result
}

func applyMask(val int, mask string) int {
	bin := fmt.Sprintf("%036b", val)
	result := make([]rune, 36)
	for i, c := range bin {
		maskChar := mask[i]
		switch maskChar {
		case '0':
			result[i] = '0'
		case '1':
			result[i] = '1'
		case 'X':
			result[i] = c
		}
	}
	return binToDec(string(result))
}

func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}

func binToDec(val string) int {
	if v, err := strconv.ParseInt(val, 2, 64); err != nil {
		panic("unable to convert")
	} else {
		return int(v)
	}
}

func readInput() string {
	b, err := ioutil.ReadFile("14/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
