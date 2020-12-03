package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1(Point{X: 3, Y: 1}))
	fmt.Printf("Part2 %d\n", part2())

}

type Point struct {
	X, Y int
}

func (p *Point) Add(p2 Point) {
	p.X += p2.X
	p.Y += p2.Y
}

func part1(move Point) int {
	var trees int
	me := Point{}
	text := readInput()
	lines := strings.Split(text, "\n")
	leng := len(lines[0])

	for i := 0; i < len(lines); {
		if string(lines[i][me.X]) == "#" {
			trees++
		}
		me.Add(move)
		if me.X > leng-1 {
			me.X = me.X % leng
		}
		i += move.Y
	}

	return trees
}
func part2() int {
	res := 1
	points := []Point{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}
	for _, p := range points {
		res = res * part1(p)
	}

	return res
}

func readInput() string {
	b, err := ioutil.ReadFile("03/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
