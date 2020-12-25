package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const TileSize = 10

func main() {
	//fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

type Q []*Tile

func (r *Q) Empty() bool  { return len(*r) == 0 }
func (r *Q) Push(p *Tile) { *r = append(*r, p) }
func (r *Q) Pop() *Tile {
	if len(*r) == 0 {
		panic("opsie, Q is empty")
	}
	el := (*r)[0]
	*r = (*r)[1:]
	return el
}

type Board struct {
	Size    int
	Data    Grid
	Monster Grid
}
type Tiles []*Tile
type Grid [][]byte
type Tile struct {
	ID                       int
	Data                     Grid
	Top, Bottom, Left, Right *Tile
}

func (r Tile) Print() {
	var result []byte
	for _, row := range r.Data {
		result = append(result, row...)
		result = append(result, '\n')
	}
	fmt.Println(string(result))
}
func (r Grid) FlipX() Grid {
	result := make([][]byte, len(r))
	for i, row := range r {
		result[i] = reverse(row)
	}
	return result
}
func (r Grid) FlipY() Grid {
	rows := len(r)
	result := make([][]byte, len(r))
	for i, row := range r {
		result[rows-i-1] = row
	}
	return result
}
func (r Grid) FlipXY() Grid {
	rows := len(r)
	result := make([][]byte, len(r))
	for i, row := range r {
		result[rows-i-1] = reverse(row)
	}
	return result
}
func (r Grid) Rotate() Grid {
	rows := len(r)
	result := make([][]byte, len(r))
	for i := 0; i < rows; i++ {
		result[i] = make([]byte, rows)
		for j := 0; j < rows; j++ {
			result[i][j] = r[rows-j-1][i]
		}
	}
	return result
}

func (r Grid) Count(i byte) int {
	result := 0
	for _, col := range r {
		for _, el := range col {
			if el == i {
				result++
			}
		}
	}

	return result
}
func (r Tile) FlipX() *Tile {
	return &Tile{ID: r.ID, Data: r.Data.FlipX(), Top: r.Top, Bottom: r.Bottom, Left: r.Right, Right: r.Left}
}
func (r Tile) FlipY() *Tile {
	return &Tile{ID: r.ID, Data: r.Data.FlipY(), Left: r.Left, Right: r.Right, Top: r.Bottom, Bottom: r.Top}
}
func (r Tile) FlipXY() *Tile {
	return &Tile{ID: r.ID, Data: r.Data.FlipXY(), Top: r.Bottom, Bottom: r.Top, Left: r.Right, Right: r.Left}
}
func (r Tile) Rotate() *Tile {
	return &Tile{ID: r.ID, Data: r.Data.Rotate(), Right: r.Top, Bottom: r.Right, Left: r.Bottom, Top: r.Left}
}

func (r Tile) GetTop() []byte    { return r.Data[0] }
func (r Tile) GetBottom() []byte { return r.Data[len(r.Data)-1] }
func (r Tile) GetLeft() []byte {
	result := make([]byte, len(r.Data))
	for i, v := range r.Data {
		result[i] = v[0]
	}
	return result
}
func (r Tile) GetRight() []byte {
	leng := len(r.Data[0]) - 1
	result := make([]byte, len(r.Data))
	for i, v := range r.Data {
		result[i] = v[leng]
	}
	return result
}

func (r *Tile) match(tile *Tile) bool {
	t := r.GetTop()
	b := r.GetBottom()
	l := r.GetLeft()
	rr := r.GetRight()

	if r.Top == nil && bytes.Equal(t, tile.GetBottom()) {
		r.Top = tile
		tile.Bottom = r
		return true
	}
	if r.Bottom == nil && bytes.Equal(b, tile.GetTop()) {
		r.Bottom = tile
		tile.Top = r
		return true
	}
	if r.Left == nil && bytes.Equal(l, tile.GetRight()) {
		r.Left = tile
		tile.Right = r
		return true
	}
	if r.Right == nil && bytes.Equal(rr, tile.GetLeft()) {
		r.Right = tile
		tile.Left = r
		return true
	}
	return false
}
func (r Tiles) MatchAll() {
	r[0] = r[0]
	r[0] = r[0].FlipY()
	q := Q{r[0]}

	visited := map[int]struct{}{}
	for !q.Empty() {
		el := q.Pop()
		for i := 0; i < len(r); i++ {
			tile := r[i]
			if r[i].ID == el.ID {
				continue
			}
			if _, ok := visited[el.ID]; ok {
				continue
			}
			if el.match(tile) {
				q.Push(tile)
			}

			r0 := tile
			r1 := r0.Rotate()
			r2 := r1.Rotate()
			r3 := r2.Rotate()
			rotations := []*Tile{r0, r1, r2, r3}

			r.transform(rotations, el, i, &q)

		}
		visited[el.ID] = struct{}{}
	}
}

func (r Tiles) transform(rotations []*Tile, el *Tile, i int, q *Q) {
	for _, rotation := range rotations {
		options := []*Tile{rotation, rotation.FlipX(), rotation.FlipY(), rotation.FlipXY()}
		for _, option := range options {
			if el.match(option) {
				r[i] = option
				q.Push(option)
				return
			}
		}
	}
}
func (r Tiles) Score() int {
	result := 1
	for _, tile := range r {
		if tile.Top == nil && tile.Left == nil {
			result *= tile.ID
		}
		if tile.Top == nil && tile.Right == nil {
			result *= tile.ID
		}
		if tile.Bottom == nil && tile.Left == nil {
			result *= tile.ID
		}
		if tile.Bottom == nil && tile.Right == nil {
			result *= tile.ID
		}
	}

	return result
}
func (r Tiles) GetTopLeft() *Tile {
	for _, tile := range r {
		if tile.Top == nil && tile.Left == nil {
			return tile
		}
	}
	panic("no tile")
}

func (r Board) Print() {
	for _, row := range r.Data {
		fmt.Print(string(row))
		fmt.Println()
	}
}
func (r *Board) Set(x, y int, tile *Tile) {
	yOffset := y * (TileSize - 2)
	xOffset := x * (TileSize - 2)

	for i, col := range tile.Data[1 : TileSize-1] {
		for j, ch := range col[1 : TileSize-1] {
			r.Data[i+yOffset][j+xOffset] = ch
		}
	}
}

func (r Board) CountMonsters() int {
	var result int
	monsterWidth := len(r.Monster[0])
	monsterHeight := len(r.Monster)

	dataWidth := len(r.Data[0])
	dataHeight := len(r.Data)

	for y := 0; y < dataHeight-monsterHeight; y++ {
		for x := 0; x <= dataWidth-monsterWidth; x++ {
			if r.matches(x, y) {
				result++
			}
		}
	}

	return result
}

func (r Board) matches(x int, y int) bool {
	for yy, row := range r.Monster {
		for xx, char := range row {
			if char != '#' {
				continue
			}
			if r.Data[y+yy][x+xx] != '#' {
				return false
			}
		}
	}
	return true
}

func ParseTile(lines []string) *Tile {
	var id int
	if _, err := fmt.Sscanf(lines[0], "Tile %d:", &id); err != nil {
		panic(err)
	}
	data := make([][]byte, len(lines)-1)
	for i, row := range lines[1:] {
		data[i] = []byte(row)
	}

	return &Tile{ID: id, Data: data}
}

func part1() int {
	tilesData := strings.Split(readInput(), "\n\n")
	tiles := make(Tiles, len(tilesData))
	for i, x := range tilesData {
		tiles[i] = ParseTile(strings.Split(x, "\n"))
	}

	tiles.MatchAll()
	return tiles.Score()
}

func part2() int {
	tilesData := strings.Split(readInput(), "\n\n")
	tiles := make(Tiles, len(tilesData))
	for i, x := range tilesData {
		tiles[i] = ParseTile(strings.Split(x, "\n"))
	}
	tiles.MatchAll()
	gridSize := int(math.Sqrt(float64(len(tiles))))
	if gridSize*gridSize != len(tiles) {
		panic("opsie")
	}
	board := createBoard(gridSize)
	start := tiles.GetTopLeft()
	x, y := 0, 0
	for start != nil {
		x = 0
		current := start
		for current != nil {
			board.Set(x, y, current)
			current = current.Right
			x++
		}
		start = start.Bottom
		y++
	}
	r0 := board.Data
	r1 := r0.Rotate()
	r2 := r1.Rotate()
	r3 := r2.Rotate()
	rotations := []Grid{r0, r1, r2, r3}
	for _, rotation := range rotations {
		options := []Grid{rotation, rotation.FlipX(), rotation.FlipY(), rotation.FlipXY()}
		for _, option := range options {
			board.Data = option
			if monsters := board.CountMonsters(); monsters != 0 {
				return board.Data.Count('#') - monsters*board.Monster.Count('#')
			}
		}
	}

	panic(":(")
}

func createBoard(size int) Board {
	data := make([][]byte, size*(TileSize-2))
	for i := range data {
		data[i] = make([]byte, size*(TileSize-2))
	}
	return Board{Data: data, Size: size, Monster: createMonster()}
}
func createMonster() Grid {
	b, err := ioutil.ReadFile("20/monster.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	result := make(Grid, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		result[i] = []byte(line)
	}
	return result
}
func readInput() string {
	b, err := ioutil.ReadFile("20/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
func reverse(data []byte) []byte {
	result := make([]byte, len(data))
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = data[j], data[i]
	}
	return result
}
