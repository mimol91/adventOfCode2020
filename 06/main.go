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

type Group struct {
	Text    string
	Answers map[rune]struct{}
}

func (g *Group) GetAnswers() {
	lines := strings.Split(g.Text, "\n")
	for _, line := range lines {
		for _, r := range line {
			g.Answers[r] = struct{}{}
		}
	}
}
func (g *Group) GetScore() int {
	result := make(map[rune]struct{})
	lines := strings.Split(g.Text, "\n")
	for _, r := range lines[0] {
		result[r] = struct{}{}
	}
	if len(lines) == 1 {
		return len(result)
	}

	for _, line := range lines[1:] {
		result = intersection(result, []rune(line))
	}

	return len(result)
}

func part1() int {
	var sum int
	lines := strings.Split(readInput(), "\n\n")
	groups := make([]Group, len(lines))

	for i, line := range lines {
		group := Group{Text: line, Answers: make(map[rune]struct{})}
		group.GetAnswers()
		groups[i] = group

		sum += len(group.Answers)
	}

	return sum
}
func part2() int {
	var sum int
	lines := strings.Split(readInput(), "\n\n")

	for _, line := range lines {
		group := Group{Text: line, Answers: make(map[rune]struct{})}
		group.GetAnswers()

		sum += group.GetScore()
	}

	return sum
}

func readInput() string {
	b, err := ioutil.ReadFile("06/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}

func intersection(runeMap map[rune]struct{}, list []rune) map[rune]struct{} {
	result := make(map[rune]struct{})
	for _, item := range list {
		if _, ok := runeMap[item]; ok {
			result[item] = struct{}{}
		}
	}

	return result
}
