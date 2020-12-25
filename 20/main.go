package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	//fmt.Printf("Part2 %d\n", part2())
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

type Tiles []*Tile
type Tile struct {
	ID                       int
	Data                     [][]byte
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
func (r Tile) FlipX() *Tile {
	data := make([][]byte, len(r.Data))
	for i, row := range r.Data {
		data[i] = reverse(row)
	}
	return &Tile{ID: r.ID, Data: data, Top: r.Top, Bottom: r.Bottom, Left: r.Right, Right: r.Left}
}
func (r Tile) FlipY() *Tile {
	rows := len(r.Data)
	data := make([][]byte, rows)
	for i, row := range r.Data {
		data[rows-i-1] = row
	}
	return &Tile{ID: r.ID, Data: data, Left: r.Left, Right: r.Right, Top: r.Bottom, Bottom: r.Top}
}
func (r Tile) FlipXY() *Tile {
	rows := len(r.Data)
	data := make([][]byte, rows)
	for i, row := range r.Data {
		data[rows-i-1] = reverse(row)
	}
	return &Tile{ID: r.ID, Data: data, Top: r.Bottom, Bottom: r.Top, Left: r.Right, Right: r.Left}
}
func (r Tile) Rotate() *Tile {
	rows := len(r.Data)
	data := make([][]byte, rows)

	for i := 0; i < rows; i++ {
		data[i] = make([]byte, rows)
		for j := 0; j < rows; j++ {
			data[i][j] = r.Data[rows-j-1][i]
		}
	}
	return &Tile{ID: r.ID, Data: data, Right: r.Top, Bottom: r.Right, Left: r.Bottom, Top: r.Left}
}

func (r Tile) Flips() [3]*Tile { return [3]*Tile{r.FlipX(), r.FlipY(), r.FlipXY()} }

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

	return 0
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
