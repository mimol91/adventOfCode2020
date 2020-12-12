package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

const emptySeat = Seat('L')
const occupiedSeat = Seat('#')
const floor = Seat('.')

type Seat rune
type NextFn func(ii, jj int) Seat

func (r Seat) isEmpty() bool    { return r == emptySeat }
func (r Seat) isOccupied() bool { return r == occupiedSeat }
func (r Seat) isFloor() bool    { return r == floor }
func (r Seat) String() string   { return fmt.Sprintf("%c", r) }

type Stats struct {
	Empty    int
	Occupied int
}
type Row []Seat
type Board struct {
	rowSize int
	Cols    []Row
}

func (r Board) Print() {
	for _, col := range r.Cols {
		for _, seat := range col {
			fmt.Print(seat)
		}
		fmt.Println()
	}
}
func (r Board) getStats(ii, jj int) Stats {
	var stats Stats
	maxI := min(ii+1, len(r.Cols)-1)
	minI := max(ii-1, 0)
	maxJ := min(jj+1, r.rowSize-1)
	minJ := max(jj-1, 0)

	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			if i == ii && j == jj {
				continue
			}

			seat := r.Cols[i][j]
			if seat.isEmpty() {
				stats.Empty++
			} else if seat.isOccupied() {
				stats.Occupied++
			}
		}
	}
	return stats
}
func (r Board) getStats2(ii, jj int) Stats {
	var stats Stats
	maxI := len(r.Cols) - 1
	maxJ := r.rowSize - 1

	directions := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, dir := range directions {
		r.generateStats(ii, jj, dir, maxI, maxJ, &stats)
	}
	return stats
}

func (r Board) generateStats(ii int, jj int, dir [2]int, maxI int, maxJ int, stats *Stats) {
	for inc := 1; ; inc++ {
		i := ii + dir[0]*inc
		j := jj + dir[1]*inc
		if i > maxI || i < 0 {
			return
		}
		if j > maxJ || j < 0 {
			return
		}
		seat := r.Cols[i][j]
		if seat == emptySeat {
			stats.Empty++
			return
		}
		if seat == occupiedSeat {
			stats.Occupied++
			return
		}
	}
}

func (r Board) getNext(ii, jj int) Seat {
	seat := r.Cols[ii][jj]
	if seat == floor {
		return floor
	}
	stats := r.getStats(ii, jj)
	if seat == emptySeat && stats.Occupied == 0 {
		return occupiedSeat
	}
	if seat == occupiedSeat && stats.Occupied >= 4 {
		return emptySeat
	}
	return seat
}

func (r Board) getNext2(ii, jj int) Seat {
	seat := r.Cols[ii][jj]
	if seat == floor {
		return floor
	}
	stats := r.getStats2(ii, jj)
	if seat == emptySeat && stats.Occupied == 0 {
		return occupiedSeat
	}
	if seat == occupiedSeat && stats.Occupied >= 5 {
		return emptySeat
	}
	return seat
}

func (r Board) Tick(fn NextFn) (Board, bool) {
	hasChange := false
	result := r.copy()
	for i := 0; i < len(r.Cols); i++ {
		for j := 0; j < r.rowSize; j++ {
			prev := result.Cols[i][j]
			result.Cols[i][j] = fn(i, j)
			if prev != result.Cols[i][j] {
				hasChange = true
			}
		}
	}
	return result, hasChange
}

func (r Board) copy() Board {
	result := Board{
		rowSize: r.rowSize,
		Cols:    make([]Row, len(r.Cols)),
	}

	for i, col := range r.Cols {
		result.Cols[i] = make(Row, result.rowSize)
		for j, char := range col {
			result.Cols[i][j] = char
		}
	}

	return result
}

func (r Board) CountSeats() int {
	var result int
	for i := 0; i < len(r.Cols); i++ {
		for j := 0; j < r.rowSize; j++ {
			if r.Cols[i][j] == occupiedSeat {
				result++
			}
		}
	}

	return result
}

func part1() int {
	text := readInput()
	board := createBoard(text)
	changed := true
	for changed {
		board, changed = board.Tick(board.getNext)
	}

	return board.CountSeats()
}

func part2() int {
	text := readInput()
	board := createBoard(text)

	changed := true
	for changed {
		board, changed = board.Tick(board.getNext2)
	}

	return board.CountSeats()
}

func createBoard(text string) Board {
	lines := strings.Split(text, "\n")
	rowSize := len(lines[0])

	board := Board{
		rowSize: rowSize,
		Cols:    make([]Row, len(lines)),
	}
	for i, line := range lines {
		board.Cols[i] = make(Row, rowSize)
		for j, char := range line {
			board.Cols[i][j] = Seat(char)
		}
	}

	return board
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readInput() string {
	b, err := ioutil.ReadFile("11/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
