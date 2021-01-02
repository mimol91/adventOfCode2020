package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Code struct {
	Key     int
	Subject int
	Loop    int
}

func (r *Code) CalculateLoop() {
	var i int
	val := 1
	for val != r.Key {
		i++
		val = transform(val, r.Subject)
	}
	r.Loop = i
}

func main() {
	fmt.Printf("Part1 %d\n", part1())
}

func part1() int {
	data := parse(readInput())
	card := data[0]
	card.CalculateLoop()

	door := data[1]
	door.CalculateLoop()

	key := getEncryptionKey(card, door)

	return key
}

func parse(input string) [2]Code {
	result := [2]Code{}
	for i, line := range strings.Split(input, "\n") {
		result[i].Key = atoi(line)
		result[i].Subject = 7
	}

	return result
}
func transform(val, sub int) int { return (val * sub) % 20201227 }
func getEncryptionKey(card Code, door Code) int {
	cardResult := 1
	doorResult := 1
	for i := 0; i < door.Loop; i++ {
		cardResult = transform(cardResult, card.Key)
	}
	for i := 0; i < card.Loop; i++ {
		doorResult = transform(doorResult, door.Key)
	}

	if cardResult == doorResult {
		return cardResult
	}
	panic("OPSIE")
}
func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}

func readInput() string {
	b, err := ioutil.ReadFile("25/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
