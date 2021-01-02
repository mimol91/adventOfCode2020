package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())

}

type Range struct {
	min, max int
}
type Section struct {
	Text  string
	Range Range
	Val   int
}

type Board struct {
	Row    Section
	Column Section
}

func (r *Range) Low()  { r.max = (r.min + r.max) / 2 }
func (r *Range) High() { r.min = (r.min + 1 + r.max) / 2 }

func (s *Section) Process(charMap map[rune]func()) {
	for _, c := range s.Text {
		charMap[c]()
	}
	if s.Range.max != s.Range.min {
		panic("unable to generate val")
	}
	s.Val = s.Range.min
}

func (Board) Parse(text string) Board {
	return Board{
		Column: Section{Text: text[7:], Range: Range{max: 7}},
		Row:    Section{Text: text[:7], Range: Range{max: 127}},
	}
}
func (b Board) Score() int { return b.Column.Val + 8*b.Row.Val }
func (b Board) String() string {
	return fmt.Sprintf("Row: %d Col: %d\n", b.Row.Val, b.Column.Val)
}
func (b *Board) Process() {
	b.Column.Process(map[rune]func(){
		'L': b.Column.Range.Low,
		'X': b.Column.Range.High,
	})
	b.Row.Process(map[rune]func(){
		'F': b.Row.Range.Low,
		'Z': b.Row.Range.High,
	})
}

func part1() int {
	var result int
	lines := strings.Split(readInput(), "\n")

	for _, line := range lines {
		board := Board{}.Parse(line)
		board.Process()
		result = max(result, board.Score())
	}

	return result
}

func part2() int {
	lines := strings.Split(readInput(), "\n")
	boards := make([]Board, len(lines))
	for i, line := range lines {
		board := Board{}.Parse(line)
		board.Process()
		boards[i] = board
	}
	sort.Slice(boards, func(i, j int) bool {
		if boards[i].Row.Val < boards[j].Row.Val {
			return true
		}
		if boards[i].Row.Val == boards[j].Row.Val {
			return boards[i].Column.Val < boards[j].Column.Val
		}
		return false
	})

	for i, board := range boards {
		prev := board
		prev.Column.Val++

		if prev.Column.Val == 8 {
			prev.Column.Val = 0
			prev.Row.Val++
		}

		if boards[i+1].Score() != prev.Score() {
			return prev.Score()
		}
	}

	panic("unable to find seatID")
}

func readInput() string {
	b, err := ioutil.ReadFile("05/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
