package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Password struct {
	min, max int
	letter   rune
	pwd      string
}

func (p Password) isValid() bool {
	found := 0
	for _, c := range p.pwd {
		if c != p.letter {
			continue
		}
		found++
		if found > p.max {
			return false
		}

	}

	return found >= p.min
}
func (p Password) isValid2() bool {
	val1 := p.pwd[p.min-1]
	val2 := p.pwd[p.max-1]
	if val1 == val2 {
		return false
	}

	return rune(val1) == p.letter || rune(val2) == p.letter
}

type Passwords []Password

func (p Passwords) validate() int {
	result := 0
	for _, pwd := range p {
		if pwd.isValid() {
			result++
		}
	}

	return result
}
func (p Passwords) validate2() int {
	result := 0
	for _, pwd := range p {
		if pwd.isValid2() {
			result++
		}
	}

	return result
}

func main() {
	fmt.Printf("Part1 %s\n", part1())
	fmt.Printf("Part2 %s\n", part2())

}

func part1() string {
	passwords := parse(readInput())

	return strconv.Itoa(passwords.validate())
}

func part2() string {
	passwords := parse(readInput())

	return strconv.Itoa(passwords.validate2())
}

func parse(input string) Passwords {
	lines := strings.Split(input, "\n")
	passwords := make(Passwords, len(lines))
	for i, line := range lines {
		var min, max int
		var letter rune
		var pwd string
		if _, err := fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &letter, &pwd); err != nil {
			panic(err)
		}

		passwords[i] = Password{min: min, max: max, letter: letter, pwd: pwd}
	}

	return passwords
}

func readInput() string {
	b, err := ioutil.ReadFile("02/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
