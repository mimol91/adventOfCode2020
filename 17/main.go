package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point3D struct {
	x, y, z int
}

func NewPoint3D(x, y, z int) Point3D { return Point3D{x: x, y: y, z: z} }
func (r Point3D) GetNeighbors() []Point {
	result := make([]Point, 0)

	for x := r.x - 1; x <= r.x+1; x++ {
		for y := r.y - 1; y <= r.y+1; y++ {
			for z := r.z - 1; z <= r.z+1; z++ {
				if x == r.x && y == r.y && z == r.z {
					continue
				}
				result = append(result, NewPoint3D(x, y, z))
			}
		}
	}

	return result
}

type Point4D struct {
	x, y, z, w int
}

func NewPoint4D(x, y, z, w int) Point4D { return Point4D{x: x, y: y, z: z, w: w} }
func (r Point4D) GetNeighbors() []Point {
	result := make([]Point, 0)

	for x := r.x - 1; x <= r.x+1; x++ {
		for y := r.y - 1; y <= r.y+1; y++ {
			for z := r.z - 1; z <= r.z+1; z++ {
				for w := r.w - 1; w <= r.w+1; w++ {
					if x == r.x && y == r.y && z == r.z && r.w == w {
						continue
					}
					result = append(result, NewPoint4D(x, y, z, w))
				}
			}
		}
	}

	return result
}

type CubeState bool

func (r CubeState) IsActive() bool { return r == true }
func (r CubeState) String() string {
	if r == true {
		return "#"
	}
	return "."
}

type Point interface {
	GetNeighbors() []Point
}
type Board map[Point]CubeState

func NewBoard() Board                             { return make(Board) }
func (r Board) SetState(p Point, state CubeState) { r[p] = state }

func (r Board) Tick() Board {
	searchSet := make(Board)
	for k, v := range r {
		searchSet[k] = v

		for _, nearby := range k.GetNeighbors() {
			if _, ok := searchSet[nearby]; !ok {
				searchSet[nearby] = false
			}
		}
	}
	result := make(Board)
	for k, val := range searchSet {
		var active int
		for _, point := range k.GetNeighbors() {
			if searchSet[point] {
				active++
			}
		}
		if val {
			result[k] = active == 3 || active == 2
		} else {
			result[k] = active == 3
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
	fmt.Printf("Part2 %d\n", part2())
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
			board.SetState(NewPoint3D(x, y, 0), state)
		}
	}
	for i := 0; i < 6; i++ {
		board = board.Tick()
	}

	return board.GetActive()
}

func part2() int {
	text := readInput()
	lines := strings.Split(text, "\n")
	board := NewBoard()
	for y, line := range lines {
		for x, char := range line {
			var state CubeState
			if char == '#' {
				state = true
			}
			board.SetState(NewPoint4D(x, y, 0, 0), state)
		}
	}
	for i := 0; i < 6; i++ {
		board = board.Tick()
	}

	return board.GetActive()
}

func readInput() string {
	b, err := ioutil.ReadFile("17/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
