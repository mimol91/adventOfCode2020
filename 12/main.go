package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Action struct {
	Val  int
	Name rune
}
type Ship struct {
	North, East int
	Direction   int
}
type Waypoint struct {
	Ship
}

func (r *Waypoint) turnRight(times int) {
	for i := 0; i < times%4; i++ {
		r.Direction = (r.Direction + 1) % 4
		east := r.North
		r.North = -r.East
		r.East = east
	}
}
func (r *Waypoint) turnLeft(times int) {
	r.turnRight(times * 3)
}
func (r *Waypoint) Execute(action Action) {
	switch action.Name {
	case 'X':
		r.turnRight(action.Val / 90)
	case 'L':
		r.turnLeft(action.Val / 90)
	default:
		r.Ship.Execute(action)
	}
}
func (r *Ship) turnRight(times int) {
	r.Direction = (r.Direction + times) % 4
}
func (r *Ship) turnLeft(times int) {
	r.turnRight(times * 3)
}
func (r *Ship) Execute(action Action) {
	switch action.Name {
	case 'N':
		r.North += action.Val
	case 'S':
		r.North -= action.Val
	case 'E':
		r.East += action.Val
	case 'W':
		r.East -= action.Val
	case 'L':
		r.turnLeft(action.Val / 90)
	case 'X':
		r.turnRight(action.Val / 90)
	case 'F':
		r.forward(action.Val)
	}
}

func (r *Ship) forward(val int) {
	switch r.Direction {
	case 0:
		r.North += val
	case 1:
		r.East += val
	case 2:
		r.North -= val
	case 3:
		r.East -= val
	}
}
func (r Ship) Distance() int {
	val := abs(r.East) + abs(r.North)
	if val < 0 {
		return -val
	}
	return val
}
func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	text := readInput()
	lines := strings.Split(text, "\n")
	ship := Ship{Direction: 1}
	for _, line := range lines {
		action := parseAction(line)
		ship.Execute(action)
	}

	return ship.Distance()
}

func parseAction(line string) Action {
	return Action{Val:  atoi(line[1:]), Name: rune(line[0]),}
}
func part2() int {
	text := readInput()
	lines := strings.Split(text, "\n")
	ship := Ship{}
	waypoint := Waypoint{Ship{North: 1, East: 10}}

	for _, line := range lines {
		action := parseAction(line)
		if action.Name == 'F' {
			ship.North += action.Val * waypoint.North
			ship.East += action.Val * waypoint.East
		} else {
			waypoint.Execute(action)
		}
	}

	return ship.Distance()
}

func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}
func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func readInput() string {
	b, err := ioutil.ReadFile("12/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
