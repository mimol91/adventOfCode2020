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

type Q struct {
	elements     []int
	preambleSize int
}

func (r *Q) Push(el int) error {
	if len(r.elements) < r.preambleSize {
		r.elements = append(r.elements, el)
		return nil
	}
	if !r.canAdd(el) {
		return fmt.Errorf("bad number %d", el)
	}

	r.elements = r.elements[1:]
	r.elements = append(r.elements, el)

	return nil
}

func (r *Q) canAdd(target int) bool {
	reamingMap := make(map[int]struct{})
	for _, el := range r.elements {
		reaming := target - el
		if _, ok := reamingMap[reaming]; ok {
			return true
		}
		reamingMap[el] = struct{}{}
	}
	return false
}

func part1() int {
	text := readInput()
	numbers := getNumbers(text)

	q := Q{preambleSize: 25}
	for _, num := range numbers {
		if err := q.Push(num); err != nil {
			return num
		}
	}

	panic("unable to find")
}

type MinMax struct {
	min int
	max int
}

func (r MinMax) Sum() int {
	return r.min + r.max
}
func (r *MinMax) Store(val int) {
	if val > r.max {
		r.max = val
	}
	if val < r.min {
		r.min = val
	}
}
func part2() int {
	text := readInput()
	numbers := getNumbers(text)
	target := part1()

	for i, num := range numbers {
		minMax := MinMax{min: math.MaxInt64, max: math.MinInt64}
		minMax.Store(num)
		sum := num
		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			minMax.Store(numbers[j])

			if sum > target {
				break
			}
			if sum == target {
				return minMax.Sum()
			}
		}
	}

	panic("unable to get")
}

func readInput() string {
	b, err := ioutil.ReadFile("09/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
func getNumbers(text string) []int {
	lines := strings.Split(text, "\n")
	result := make([]int, len(lines))
	for i, line := range lines {
		result[i] = atoi(line)
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
