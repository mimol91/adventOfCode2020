package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Cards []int

func (r Cards) Size() int   { return len(r) }
func (r Cards) Empty() bool { return r.Size() == 0 }
func (r *Cards) Push(v int) { *r = append(*r, v) }
func (r *Cards) Pop() int {
	if len(*r) == 0 {
		panic("opsie, Q is empty")
	}
	el := (*r)[0]
	*r = (*r)[1:]
	return el
}
func (r *Cards) Win(a, b int) {
	r.Push(a)
	r.Push(b)
}
func (r Cards) CanPlaySubGame(card int) bool {
	return r.Size() >= card
}

func (r Cards) Print() {
	fmt.Println(r)
}

func (r Cards) GetScore() (result int) {
	for i := len(r); i > 0; i-- {
		result += i * r.Pop()
	}
	return result
}

func (r Cards) Slice(len int) Cards {
	return r[:len:len]
}

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	players := strings.Split(readInput(), "\n\n")
	player1 := toCards(players[0])
	player2 := toCards(players[1])
	for !player1.Empty() && !player2.Empty() {
		p1Card := player1.Pop()
		p2Card := player2.Pop()
		if p1Card > p2Card {
			player1.Win(p1Card, p2Card)
		} else {
			player2.Win(p2Card, p1Card)
		}
	}

	return player1.GetScore() + player2.GetScore()
}

func toCards(text string) Cards {
	vals := strings.Split(text, "\n")[1:]
	result := make([]int, len(vals))
	for i, c := range vals {
		result[i] = atoi(c)
	}
	return result

}

func part2() int {
	players := strings.Split(readInput(), "\n\n")
	player1 := toCards(players[0])
	player2 := toCards(players[1])

	if play(&player1, &player2, 0) {
		return player1.GetScore()
	}

	return player2.GetScore()
}
func play(player1, player2 *Cards, nestLvl int) bool {
	history := make(map[string]struct{})

	for !player1.Empty() && !player2.Empty() {
		if _, ok := history[fmt.Sprintf("%v", player1)]; ok {
			return true
		}
		if _, ok := history[fmt.Sprintf("%v", player2)]; ok {
			return true
		}

		history[fmt.Sprintf("%v", player1)] = struct{}{}
		history[fmt.Sprintf("%v", player2)] = struct{}{}

		p1Card := player1.Pop()
		p2Card := player2.Pop()
		if !player1.CanPlaySubGame(p1Card) || !player2.CanPlaySubGame(p2Card) {
			if p1Card > p2Card {
				player1.Win(p1Card, p2Card)
			} else {
				player2.Win(p2Card, p1Card)
			}
			continue
		}
		p1Copy := player1.Slice(p1Card)
		p2Copy := player2.Slice(p2Card)

		if play(&p1Copy, &p2Copy, nestLvl+1) {
			player1.Win(p1Card, p2Card)
		} else {
			player2.Win(p2Card, p1Card)
		}
	}

	if player1.Empty() {
		return false
	}
	return true
}

func readInput() string {
	b, err := ioutil.ReadFile("22/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}
