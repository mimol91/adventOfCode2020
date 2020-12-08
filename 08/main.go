package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const maxIterations = 9999

type Registers [2]int //0 Acc, 1 IP
type Instruction func(registers Registers, val int) Registers

var Instructions = map[string]Instruction{
	"nop": func(registers Registers, _ int) Registers { return registers },
	"jmp": func(registers Registers, val int) Registers { registers[1] += val - 1; return registers },
	"acc": func(registers Registers, val int) Registers { registers[0] += val; return registers },
}

type Operation struct {
	Num         int
	Name        string
	Instruction Instruction
	Val         int
}

func (r Operation) Execute(registers Registers) Registers {
	return r.Instruction(registers, r.Val)
}

func main() {
	fmt.Printf("Part1 %d\n", part1())
	fmt.Printf("Part2 %d\n", part2())
}

func part1() int {
	registers := Registers{}
	executedInstructions := map[int]struct{}{}
	text := strings.Split(readInput(), "\n")
	operations := getOperations(text)
	for i := 0; i < maxIterations; i++ {
		operation := operations[registers[1]]
		if _, ok := executedInstructions[operation.Num]; ok {
			return registers[0]
		}
		registers = operation.Execute(registers)
		executedInstructions[operation.Num] = struct{}{}
		registers[1]++
	}

	panic("Unable to get correct val")
}

func part2() int {
	text := strings.Split(readInput(), "\n")
	operations := getOperations(text)
	for i := 0; i < len(operations); i++ {
		registers := Registers{}
		operationsCopy := make([]Operation, len(operations))
		copy(operationsCopy, operations)

		if operationsCopy[i].Name == "acc" {
			continue
		}

		if operationsCopy[i].Name == "nop" {
			operationsCopy[i].Name = "jmp"
			operationsCopy[i].Instruction = Instructions["jmp"]
		} else if operationsCopy[i].Name == "jmp" {
			operationsCopy[i].Name = "nop"
			operationsCopy[i].Instruction = Instructions["nop"]
		}
		if val, err := execute(operationsCopy, registers); err == nil {
			return val
		}
	}

	panic("Unable to get correct val")
}

func execute(operations []Operation, registers Registers) (int, error) {
	maxLen := len(operations)
	executedInstructions := map[int]struct{}{}

	for i := 0; i < maxIterations; i++ {
		if registers[1] >= maxLen {
			return registers[0], nil
		}
		operation := operations[registers[1]]
		if _, ok := executedInstructions[operation.Num]; ok {
			return 0, fmt.Errorf("already executed")
		}
		registers = operation.Execute(registers)
		registers[1]++
	}
	return 0, fmt.Errorf("unable to terminame")
}

func getOperations(text []string) []Operation {
	operations := make([]Operation, len(text))
	for i, line := range text {
		var name, stringVal string
		if _, err := fmt.Sscanf(line, "%s %s", &name, &stringVal); err != nil {
			panic(err)
		}
		operation := Operation{
			Num:         i,
			Name:        name,
			Instruction: Instructions[name],
			Val:         atoi(stringVal),
		}
		operations[i] = operation
	}
	return operations
}

func readInput() string {
	b, err := ioutil.ReadFile("08/input.txt")
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
