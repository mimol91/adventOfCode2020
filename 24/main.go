package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`(e)|(se)|(sw)|(w)|(nw)|(ne)`)

type Hex struct {
	X, Y, Z int
}

func (r Hex) E() Hex  { r.X += 1; r.Y += -1; r.Z += 0; return r }
func (r Hex) NE() Hex { r.X += 1; r.Y += 0; r.Z += -1; return r }
func (r Hex) SE() Hex { r.X += 0; r.Y += -1; r.Z += 1; return r }

func (r Hex) W() Hex  { r.X += -1; r.Y += 1; r.Z += 0; return r }
func (r Hex) NW() Hex { r.X += 0; r.Y += 1; r.Z += -1; return r }
func (r Hex) SW() Hex { r.X += -1; r.Y += 0; r.Z += 1; return r }

func (r Hex) ActionMap() map[string]func(r Hex) Hex {
	return map[string]func(r Hex) Hex{
		"e":  Hex.E,
		"ne": Hex.NE,
		"se": Hex.SE,
		"w":  Hex.W,
		"nw": Hex.NW,
		"sw": Hex.SW,
	}
}

func (r Hex) Parse(text string) Hex {
	actionMap := r.ActionMap()
	matches := pattern.FindAllString(text, -1)
	for _, m := range matches {
		r = actionMap[m](r)
	}

	return r
}

type Neighbor struct {
	Hex     Hex
	IsBlack bool
}

func (r Hex) getNeighbours(tiles map[Hex]struct{}) [6]Neighbor {
	var result [6]Neighbor
	var i int
	for _, fn := range r.ActionMap() {
		hex := fn(r)
		_, ok := tiles[hex]
		result[i] = Neighbor{Hex: hex, IsBlack: ok}
		i++
	}

	return result
}

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	tiles := make(map[Hex]struct{})
	lines := strings.Split(readInput(), "\n")

	for _, line := range lines {
		var p Hex
		p = p.Parse(line)
		if _, ok := tiles[p]; ok {
			delete(tiles, p)
		} else {
			tiles[p] = struct{}{}
		}
	}

	return len(tiles)
}
func part2() int {
	tiles := make(map[Hex]struct{})
	lines := strings.Split(readInput(), "\n")

	for _, line := range lines {
		var p Hex
		p = p.Parse(line)
		if _, ok := tiles[p]; ok {
			delete(tiles, p)
		} else {
			tiles[p] = struct{}{}
		}
	}

	for i := 1; i <= 100; i++ {
		flipTiles(tiles)
	}

	return len(tiles)
}

func flipTiles(tiles map[Hex]struct{}) {
	toAdd := make([]Hex, 0)
	toRemove := make([]Hex, 0)
	whiteTiles := make(map[Hex]struct{})

	for tile := range tiles {
		var blackCount int
		neighbours := tile.getNeighbours(tiles)
		for _, neighbour := range neighbours {
			if neighbour.IsBlack {
				blackCount++
			} else {
				whiteTiles[neighbour.Hex] = struct{}{}
			}
		}
		if blackCount == 0 || blackCount > 2 {
			toRemove = append(toRemove, tile)
		}
	}
	for whiteTile := range whiteTiles {
		var blackCount int
		neighbours := whiteTile.getNeighbours(tiles)
		for _, neighbour := range neighbours {
			if neighbour.IsBlack {
				blackCount++
			}
		}
		if blackCount == 2 {
			toAdd = append(toAdd, whiteTile)
		}
	}

	for _, v := range toAdd {
		tiles[v] = struct{}{}
	}
	for _, v := range toRemove {
		if _, ok := tiles[v]; ok {
			delete(tiles, v)
		}
	}
}

func readInput() string {
	b, err := ioutil.ReadFile("24/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
