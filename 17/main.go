package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

func NewPoint(x, y, z int) Point { return Point{x: x, y: y, z: z} }
func (r Point) String() string   { return fmt.Sprintf("%d.%d.%d", r.x, r.y, r.z) }
func (r Point) getNeighbors() []Point {
	result := make([]Point, 0)

	for x := r.x - 1; x <= r.x+1; x++ {
		for y := r.y - 1; y <= r.y+1; y++ {
			for z := r.z - 1; z <= r.z+1; z++ {
				if x == r.x && y == r.y && z == r.z {
					continue
				}
				result = append(result, Point{x: x, y: y, z: z})
			}
		}
	}

	return result
}

func (r Point) FromString(k string) Point {
	el := strings.Split(k, ".")
	return NewPoint(atoi(el[0]), atoi(el[1]), atoi(el[2]))
}

type CubeState bool

func (r CubeState) IsActive() bool { return r == true }
func (r CubeState) String() string {
	if r == true {
		return "#"
	}
	return "."
}

type Board map[string]CubeState

func NewBoard() Board                             { return make(Board) }
func (r Board) SetState(p Point, state CubeState) { r[p.String()] = state }

func (r Board) Tick() Board {
	searchSet := make(Board)
	for k, v := range r {
		searchSet[k] = v
		p := Point{}.FromString(k)

		for _, nearby := range p.getNeighbors() {
			if _, ok := searchSet[nearby.String()]; !ok {
				searchSet[nearby.String()] = false
			}
		}
	}
	result := make(Board)
	for k, val := range searchSet {
		var active int
		p := Point{}.FromString(k)
		for _, point := range p.getNeighbors() {
			if searchSet[point.String()] {
				active++
			}
		}
		if val {
			result[p.String()] = active == 3 || active == 2
		} else {
			result[p.String()] = active == 3
		}
	}

	return result
}

func (r Board) GetActive() int {
	var result int

	for _, v := range r {
		if v.IsActive() {
			result++
		}
	}
	return result
}

func main() {
	fmt.Printf("Part1 %d\n", part1())
	//fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	text := readInput()
	lines := strings.Split(text, "\n")
	board := NewBoard()
	for y, line := range lines {
		for x, char := range line {
			var state CubeState
			if char == '#' {
				state = true
			}
			board.SetState(Point{x: x, y: y, z: 0}, state)
		}
	}

	board = board.Tick()
	board = board.Tick()
	board = board.Tick()
	board = board.Tick()
	board = board.Tick()
	board = board.Tick()

	return board.GetActive()
}
func part2() int {

	return 0
}

func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}
func readInput() string {
	b, err := ioutil.ReadFile("17/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
