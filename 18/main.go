package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part1 %d\n", part1())
	//fmt.Printf("Part2 %d\n", part2())
}

type Operator rune

func (r Operator) Execute(a, b int) int {
	if r == '+' {
		return a + b
	}
	return a * b
}

type Calculator struct {
	Operator Operator
	A        int
}

func (r Calculator) Result() int { return r.A }
func (r *Calculator) Store(val string) {
	if val == "+" || val == "*" {
		r.Operator = Operator(rune(val[0]))
		return
	}
	nr := atoi(val)
	if r.Operator == 0 {
		r.A = nr
		return
	}

	r.A = r.Operator.Execute(r.A, nr)
	r.Operator = 0
}
func Calc(text string) int {
	calc := Calculator{}
	parts := strings.Split(text, " ")
	for _, val := range parts {
		calc.Store(val)
	}
	return calc.Result()
}

func ParseLine(text string) int {
	for {
		openB, closeB := getBrackets(text)
		if closeB == 0 {
			break
		}
		inner := text[openB+1 : closeB]
		text = text[0:openB] + strconv.Itoa(Calc(inner)) + text[closeB+1:]
	}
	return Calc(text)
}

func part1() int {
	var result int
	text := readInput()
	for _, line := range strings.Split(text, "\n") {
		result += ParseLine(line)
	}

	return result
}
func getBrackets(text string) (int, int) {
	closingBrackets := len(text)
	openingBrackets := 0
	for i := closingBrackets - 1; i >= 0; i-- {
		if text[i] == ')' {
			closingBrackets = i
		}
	}
	if closingBrackets == len(text) {
		return 0, 0
	}
	for i := closingBrackets - 1; i >= 0; i-- {
		if text[i] == '(' {
			openingBrackets = i
			break
		}
	}

	return openingBrackets, closingBrackets
}

func part2() int {

	return 0
}

func atoi(val string) int {
	if res, err := strconv.Atoi(val); err != nil {
		panic("unable to convert")
	} else {
		return res
	}
}

func readInput() string {
	b, err := ioutil.ReadFile("18/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
